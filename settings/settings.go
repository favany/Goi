package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存配置的所有信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name             string `mapstructure:"name"`
	Mode             string `mapstructure:"mode"`
	Version          string `mapstructure:"version"`
	Port             int    `mapstructure:"port"`
	*LogConfig       `mapstructure:"log"`
	*MySQLConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
	*SnowFlakeConfig `mapstructure:"snowflake"`
}

type LogConfig struct {
	Level      string `mapstructure:"Level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackUps int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	DB       string `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type SnowFlakeConfig struct {
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

func Init(fileName string) (err error) {
	// 方法1：直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	//viper.SetConfigFile("./conf/config.yaml")
	// 绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("/Users/admin/Projects/Web/config.yaml")

	// 方法二：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	// 配置文件名不需要带后缀
	// 配置文件位置可配置多个
	//viper.SetConfigName("config") // 指定配置文件名称（不需要带后缀）
	//viper.AddConfigPath(".")      // 指定查找配置文件的路径（这里指相对路径）

	viper.SetConfigFile(fileName)

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置文件失败
		fmt.Printf("read config file failed, err:%v\n", err)
		return err
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		// 把读取到的配置信息反序列化到 Conf 变量中
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
			return
		}
	})
	return
}
