package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var Conf Config

type Config struct {
	Http  Http  `yaml:"http"`
	Kafka Kafka `yaml:"kafka"`
	Redis Redis `yaml:"redis"`
	Mongo Mongo `yaml:"mongo"`
}

type Http struct {
	Name                     string        `yaml:"name"`
	Address                  string        `yaml:"address"`
	Version                  string        `yaml:"version"`
	ReadTimeout              time.Duration `yaml:"readTimeout"`
	WriteTimeout             time.Duration `yaml:"writeTimeout"`
	MaxConnPerHost           int           `yaml:"maxConnPerHost"`
	MaxIdleConnDuration      time.Duration
	MaxConnWaitTimeout       time.Duration
	NoDefaultUserAgentHeader bool `yaml:"noDefaultUserAgentHeader"`
}

type Kafka struct {
	Address string `yaml:"address"`
	TopIc   string `yaml:"topIc"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int64  `yaml:"DB"`
}

type Mongo struct {
	User             string `yaml:"user"`
	Password         string `yaml:"password"`
	Address          string `yaml:"address"`
	DB               string `yaml:"db"`
	StressDebugTable string `yaml:"stressDebugTable"`
	SceneDebugTable  string `yaml:"sceneDebugTable"`
	ApiDebugTable    string `yaml:"apiDebugTable"`
}

func InitConfig() {

	var conf string
	flag.StringVar(&conf, "c", "./dev.yaml", "配置文件,默认为conf文件夹下的dev文件")
	if !flag.Parsed() {
		flag.Parse()
	}

	viper.SetConfigFile(conf)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err = viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal error config file: %w", err))
	}

	fmt.Println("config initialized")

}
