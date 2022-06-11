package setting

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var lg *zap.Logger


//初始化zapLogger
func InitZapLogger(cfg *LoggerConfig, mode string) (err error) {
	encoder := getEncoder()
	writerSyncer := getLogWriter(cfg)
	level := new(zapcore.Level)
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}

	var core zapcore.Core
	if mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writerSyncer, level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
			)
	} else {
		core = zapcore.NewCore(encoder, writerSyncer, level)
	}

	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) //替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.L().Info("init logger success")
	return nil
}

//zapcore.Field: 一组键值对参数
//编码器，解决如何写入日志
func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}

//指定日志将写到哪里
//添加日志切割归档功能，使用第三方库Lumberjack来实现
func getLogWriter(cfg *LoggerConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: cfg.Filename,
		MaxSize: cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge: cfg.MaxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}


//接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method ),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("IP", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),

			)
	}
}

//recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			var nil any
			if err := recover(); err != nil {
				var brokenPipe bool
				var err any
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer"){
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(
						c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						)
					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recoery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),

						)
				} else {
					zap.L().Error("[Recoery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}




}
