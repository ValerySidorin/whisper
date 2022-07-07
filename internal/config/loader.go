package config

import (
	"path"
	"strings"

	"github.com/spf13/viper"
)

var gConfig Configuration

func Register() *Configuration {
	return &gConfig
}

func LoadDefaultConfigByViper(cfgPath string) (Configuration, error) {
	viper := viper.New()
	var base = path.Base(cfgPath)
	var ext = path.Ext(base)
	var name = strings.TrimSuffix(base, ext)
	viper.SetConfigName(name)
	viper.SetConfigType(ext[1:])
	viper.AddConfigPath(strings.TrimSuffix(cfgPath, base))
	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, err
	}
	var cfg Configuration
	if err := viper.Unmarshal(&cfg); err != nil {
		return Configuration{}, err
	}
	gConfig = cfg
	return gConfig, nil
}
