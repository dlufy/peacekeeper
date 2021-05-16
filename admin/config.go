package admin

import (
	"fmt"

	gcfg "gopkg.in/gcfg.v1"
)

type MainConfig struct {
	Server struct {
		Port string
	}
	Database struct {
		Password string
		DBName   string
	}
	AWS struct {
		SecretAcessKey string
		AcessKey       string
		Password       string
		UserName       string
		BucketName     string
		Region         string
	}
}

var MainConfigValue *MainConfig

func Initialize(filePath string) {
	MainConfigValue = &MainConfig{}
	err := gcfg.ReadFileInto(MainConfigValue, filePath)
	if err != nil {
		fmt.Println("[config]error while initializing the config", err)
		return
	}

}

func GetConfig() *MainConfig {

	return MainConfigValue
}
