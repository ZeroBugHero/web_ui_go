package configs

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"web_ui_go/models"
)

var config *models.Config

func ConfigLoad() *models.Config {
	readFile, err := os.ReadFile("../data/config/conf.yml")
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	err = yaml.Unmarshal(readFile, &config)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	return config
}
