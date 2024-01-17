package server

import (
	"os"

	"github.com/starfork/stargo/config"
	"gopkg.in/yaml.v3"
)

// 当前服务配置
type ServiceConfig struct {
	Server *config.Config
}

func LoadConfig(f string) *ServiceConfig {
	conf := &ServiceConfig{}
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&conf); err != nil {
		panic(err)
	}
	return conf
}
