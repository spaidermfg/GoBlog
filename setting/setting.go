package setting

//读取配置文件

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	*AppConfig    `mapstructure:"app"`
	*LoggerConfig `mapstructure:"log"`
	*MysqlConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Port int64  `mapstructure:"port"`
}

type LoggerConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructrue:"max-size"`
	MaxBackups int    `mapstructrue:"max-backups"`
	MaxAge     int    `mapstructrue:"max-age"`
}

type MysqlConfig struct {
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	MaxActive   int    `mapstructure:"maxactive"`
	MaxIdle     int    `mapstructure:"maxidle"`
	IdleTimeout int    `mapstructure:"idletimeout"`

}

func InitViper() (err error) {

	viper.SetConfigFile("./conf/config.yaml")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf/")
	//viper.SetConfigFile(filepath)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件信息读取错误", err)
		return err
	}
	//将读取的配置文件保存至全局变量Conf
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Println("保存至全局变量失败", err)
		return err
	}

	//监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Println("更新全局变量失败")
		}
	})

	fmt.Println("save config file success...")
	return nil
}
