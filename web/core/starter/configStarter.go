package starter

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"learn-go/web/core"
	"learn-go/web/core/context"
	"log"
	"path"
	"path/filepath"
	"strings"
)

type ConfigStarter struct {
	ConfigPath string
	AbstractStarter
}

const (
	GlobalConfigKey = "SystemConfig"
)

func (config *ConfigStarter) Init(context context.ApplicationContext) {
	configPath := "./resource"
	v := viper.New()
	if config.ConfigPath != "" && len(config.ConfigPath) > 0 {
		if strings.Contains(config.ConfigPath, string(filepath.Separator)) {
			v.SetConfigFile(config.ConfigPath)
		} else {
			v.AddConfigPath(configPath)
			v.SetConfigName(config.ConfigPath)
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
					logrus.Errorf("config %s type can not parse", file.Name())
				}
			}
		}
	}
	if err := v.ReadInConfig(); err != nil {
		panic("read config failed error message:" + err.Error())
	}
	context.Set(GlobalConfigKey, *v)
	logrus.Info("config init success")
}

func (config *ConfigStarter) Finalize(context context.ApplicationContext) {
	context.Remove(GlobalConfigKey)
}

func (config *ConfigStarter) GetOrder() int {
	return core.Int32Min + 1
}
