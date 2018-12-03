package main

import (
	"fmt"
	"github.com/spf13/viper"
	"qdk/qdk.v1/http"
)

func main() {
	viper.SetConfigName("dev")
	viper.AddConfigPath("./qdk.v1/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	MysqlUri := viper.GetString("mysql_uri")
	TsdbUri := viper.GetString("tsdb_uri")
	HttpAddr := viper.GetString("http_addr")

	cfg := &http.Config{
		MysqlUri: MysqlUri,
		TsdbUri:  TsdbUri,
		HttpAddr: HttpAddr,
	}

	server := http.NewServer(cfg)
	server.Start()
}
