package main

import (
	"GoBlog/database"
	"GoBlog/setting"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)



var db *sqlx.DB


func main() {
	//getFunc()
	//postFunc()
	//uploadFile()
	//uploadSelectFile()
	//uplaodManyFile()
	//error()

	//配置管理工具
	if err := setting.InitViper(); err != nil {
		fmt.Errorf("读取配置文件初始化失败", err)
		return
	}

	//日志管理工具
	if err := setting.InitZapLogger(setting.Conf.LoggerConfig, setting.Conf.Mode); err != nil {
		fmt.Println("init logger failed",err)
		return
	}
	defer zap.L().Sync() //将缓存中的日志同步到日志文件中

	//配置数据库
	if err := database.InitDB(setting.Conf.MysqlConfig ); err != nil {
		fmt.Println("init mysql failed", err)
		return
	}
	//database.Close()

}



func useZap() {
	r := gin.Default()
	// 注册zap相关中间件
	r.Use(setting.GinLogger(), setting.GinRecovery(true))

	r.GET("/hello", func(c *gin.Context) {
		// 假设你有一些数据需要记录到日志中
		var (
			name = "q1mi"
			age  = 18
		)
		// 记录日志并使用zap.Xxx(key, val)记录相关字段
		zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))

		c.String(http.StatusOK, "hello liwenzhou.com!")
	})

	addr := fmt.Sprintf(":%v", setting.Conf.AppConfig.Port)
	r.Run(addr)
}


/*func selectOne() {
	var db *sqlx.DB
	sqlStr := "select * from employees where order = ?"
	var o order
	err := db.Get(&o, sqlStr, 1)
	if err != nil {
		zap.L().Error("select employees where order = 1 fail")
		return
	}
	fmt.Printf("order_id: %v, order_name: %v\n", o.orderId, o.orderName)
}*/

func getFunc() {
	//创建默认路由
	r := gin.Default()
	//绑定路由规则和执行函数，当访问/ping接口时就会执行后面的函数
	r.GET("/ping", func(c *gin.Context) {

		//该函数返回json字符串
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", setting.Conf.AppConfig.Port)) // 监听并在 0.0.0.0:8080 上启动服务
}

func postFunc() {

	r := gin.Default()
	/*
		post常见传输格式
		application/json
		application/x-www-form-urlencoded
		application/xml
		multipart/form-data
	*/
	r.POST("/post.do", func(c *gin.Context) {
		message := c.DefaultPostForm("message", "default message")
		name := c.PostForm("name")
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"name":    name,
		})
	})
	r.Run()
}

func uploadFile() {
	r := gin.Default()

	r.POST("/upload.do", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
		}
		c.SaveUploadedFile(file, file.Filename)
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"filename": file.Filename,
		})
	})
	r.Run()
}

func uploadSelectFile() {
	r := gin.Default()

	r.POST("/uploadimg.do", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
		}

		if file.Size > 1024*1024*2 {
			c.JSON(http.StatusOK, gin.H{
				"message": "文件大小不能超过2MB",
			})
			return
		}

		if file.Header.Get("Content-type") != "iamge/png" {
			c.JSON(http.StatusOK, gin.H{
				"message": "只能上传iamge或png格式",
			})
			return
		}

		c.SaveUploadedFile(file, file.Filename)
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"filename": file.Filename,
		})
	})
	r.Run()
}

func uplaodManyFile() {
	r := gin.Default()

	r.POST("/uploadMultFile.do", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
			return
		}

		files := form.File["files"]
		for _, file := range files {
			err := c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "文件上传失败",
					"err":     err,
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "ok",
			"filecount": len(files),
		})
	})
	r.Run()
}

func error() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "页面丢了",
		})
	})
	r.Run()
}
