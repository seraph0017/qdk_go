package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"qdk/qdk.v1/http"
)

func main() {
	if os.Getenv(gin.ENV_GIN_MODE) == "release" {
		viper.SetConfigName("prod")
	} else {
		viper.SetConfigName("dev")
	}
	viper.AddConfigPath("./qdk.v1/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	MysqlUri := viper.GetString("mysql_uri")
	TsdbUri := viper.GetString("tsdb_uri")
	HttpAddr := viper.GetString("http_addr")
	LogPath := viper.GetString("log_path")

	cfg := &http.Config{
		MysqlUri: MysqlUri,
		TsdbUri:  TsdbUri,
		HttpAddr: HttpAddr,
		LogPath:  LogPath,
	}

	server := http.NewServer(cfg)
	server.Start()
}
