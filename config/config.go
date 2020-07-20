package config 

import (
	"github.com/eduardogpg/gonv"
	"fmt"
)

type DatabaseConfig struct{
	Username string
	Password string
	Host string
	Port int
	Database string
}

var database *DatabaseConfig

func init(){
	database = &DatabaseConfig{}
	database.Username = gonv.GetStringEnv("USERNAME","root")
	database.Password = gonv.GetStringEnv("PASSWORD","")
	database.Host = gonv.GetStringEnv("HOST","localhost")
	database.Port = gonv.GetIntEnv("PORT",3306)
	database.Database = gonv.GetStringEnv("DATABASE","project_go_web")
}

func GetUrlDatabase() string{
	return database.url()
}

func (this *DatabaseConfig) url() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",this.Username, this.Password, this.Host, this.Port, this.Database)
}
