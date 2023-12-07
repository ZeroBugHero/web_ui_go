package configs

import (
	"github.com/ZeroBugHero/web_ui_go/models"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var config *models.Config

func ConfigLoad() *models.Config {
	readFile, err := os.ReadFile("/Users/pizazz/Desktop/learn/code/web_ui/web_ui_go/data/config/conf.yml")
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	err = yaml.Unmarshal(readFile, &config)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	return config
}
