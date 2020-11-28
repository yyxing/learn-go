package base

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"learn-go/web/core"
	"log"
	"path"
	"strings"
)

type ConfigStarter struct {
	configName string
	AbstractStarter
}

const (
	GlobalConfigKey = "SystemConfig"
)

func (config *ConfigStarter) Init(context core.ApplicationContext) {
	configPath := "./resource"
	v := viper.New()
	if config.configName != "" && len(config.configName) > 0 {
		if strings.Contains(config.configName, "/") {
			v.SetConfigFile(config.configName)
		} else {
			v.AddConfigPath(configPath)
			v.SetConfigName(config.configName)
		}
	} else {
		files, err := ioutil.ReadDir(configPath)
		if err != nil {
			log.Println("Find config file failed use default config")
		}
		v.AddConfigPath(configPath)
		suffix := ""
		for _, file := range files {
			if !file.IsDir() {
				suffix = path.Ext(file.Name())[1:]
				switch suffix {
				case "yml", "yaml", "properties", "ini":
					v.SetConfigName(file.Name())
					v.SetConfigType(suffix)
				default:
					log.Println("config type can not parse")
				}
			}
		}
	}
	if err := v.ReadInConfig(); err != nil {
		log.Println("read config failed error message:", err)
		return
	}
	context.Set(GlobalConfigKey, *v)
	log.Println("config init success")
}

func (config *ConfigStarter) Finalize(context core.ApplicationContext) {
	context.Remove(GlobalConfigKey)
}

func (config *ConfigStarter) GetOrder() int {
	return ^int(^uint32(0) >> 1)
}
