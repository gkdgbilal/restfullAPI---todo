package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var C appConf

func Init(path ...string) {
	if len(path) == 0 {
		path = []string{"./configs/"}
	}

	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	viper.Set("env", env)

	viper.SetConfigName(fmt.Sprintf("%s.%s", "application", env)) // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path[0]) // config file path
	viper.AutomaticEnv()         // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("fatal error config file: default \n", err)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatal("fatal error config file: default \n", err)
	}
	log.Println("Config initialize!")
}

func DBConnectionURL() string {
	return fmt.Sprintf("mongodb://%s:%s", C.Database.Host, C.Database.Port)
}

type appConf struct {
	Env string
	App struct {
		Name string
		Port string
	}
	Security struct {
		Jwt struct {
			Exp string // Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
			Key string
		}
	}
	Database struct {
		Host     string
		Port     string
		Name     string
		Username string
		Password string
	}
	Minio struct {
		Host   string
		Port   string
		Region string
		Access string
		Secret string
	}
}
