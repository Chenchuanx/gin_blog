package core

import (
	"fmt"
	"goBlog/config"
	"goBlog/global"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// 读取配置
func InitConf() {
	// ReadFile
	const ConfigFile = "setting.yaml"
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	// 解析yaml
	c := &config.Config{}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
	// fmt.Println(c)
}
