package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"qdk/qdk.v1/middleware"
	"qdk/qdk.v1/misc"
	"qdk/qdk.v1/service"
)

type (
	Server struct {
		App *gin.Engine

		HttpAddr string

		TsdbUri  string
		MysqlUri string

		PiMgr   service.PiMgr
		KengMgr service.KengMgr

		Log *logrus.Logger
	}

	Config struct {
		TsdbUri  string
		MysqlUri string
		HttpAddr string
		LogPath  string
	}
)

func (s *Server) Start() {
	s.App.Run(s.HttpAddr)
}

func NewServer(cfg *Config) *Server {
	r := gin.Default()
	Log := misc.NewLogger(cfg.LogPath)

	PiMgr, err := service.NewPiMgr(&service.PiConfig{MysqlUri: cfg.MysqlUri, TsdbUri: cfg.TsdbUri, Log: Log})
	if err != nil {
		Log.Error(err)
	}

	KengMgr, err := service.NewKengMgr(&service.KengConfig{MysqlUri: cfg.MysqlUri, TsdbUri: cfg.TsdbUri, Log: Log})
	if err != nil {
		Log.Error(err)
	}

	s := &Server{
		App: r,

		HttpAddr: cfg.HttpAddr,
		TsdbUri:  cfg.TsdbUri,
		MysqlUri: cfg.MysqlUri,

		PiMgr:   PiMgr,
		KengMgr: KengMgr,

		Log: Log,
	}

	pi := s.App.Group("/pi")
	pi.Use(middleware.DeviceValidate(PiMgr))
	pi.POST("/ultrasonic", s.PostUltraSonicHandler)

	keng := s.App.Group("/keng")
	keng.GET("/", s.KengIndexHandler)

	return s
}
