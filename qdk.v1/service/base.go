package service

import (
	"github.com/bluebreezecf/opentsdb-goclient/client"
	"github.com/jinzhu/gorm"
)

type Base struct {
	Conn    *gorm.DB
	TsdbCli client.Client
}
