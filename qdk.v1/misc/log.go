package misc

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

func NewLogger(path string) *logrus.Logger {

	writerInfo, err := rotatelogs.New(
		path+"info.log."+time.Now().Format("20060102"),
		rotatelogs.WithLinkName(path+"info.log"),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	writerError, err := rotatelogs.New(
		path+"error.log."+time.Now().Format("20060102"),
		rotatelogs.WithLinkName(path+"error.log"),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	if err != nil {
		panic(err)
	}

	log := logrus.New()
	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writerInfo,
			logrus.ErrorLevel: writerError,
		},
		&logrus.JSONFormatter{},
	))
	return log

}
