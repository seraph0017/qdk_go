package service

import (
	"github.com/bluebreezecf/opentsdb-goclient/client"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Base struct {
	Conn    *gorm.DB
	TsdbCli client.Client
	Log     *logrus.Logger
}
