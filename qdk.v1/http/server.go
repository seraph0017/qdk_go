package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qdk/qdk.v1/middleware"
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
	}

	Config struct {
		TsdbUri  string
		MysqlUri string
		HttpAddr string
	}
)

func (s *Server) Start() {
	s.App.Run(s.HttpAddr)
}

func NewServer(cfg *Config) *Server {
	r := gin.Default()

	PiMgr, err := service.NewPiMgr(&service.PiConfig{MysqlUri: cfg.MysqlUri, TsdbUri: cfg.TsdbUri})
	if err != nil {
		fmt.Println("err", err)
	}

	KengMgr, err := service.NewKengMgr(&service.KengConfig{MysqlUri: cfg.MysqlUri, TsdbUri: cfg.TsdbUri})
	if err != nil {
		fmt.Println("err", err)
	}

	s := &Server{
		App: r,

		HttpAddr: cfg.HttpAddr,
		TsdbUri:  cfg.TsdbUri,
		MysqlUri: cfg.MysqlUri,

		PiMgr:   PiMgr,
		KengMgr: KengMgr,
	}

	pi := s.App.Group("/pi")
	pi.Use(middleware.DeviceValidate(PiMgr))
	pi.POST("/ultrasonic", s.PostUltraSonicHandler)

	keng := s.App.Group("/keng")
	keng.GET("/", s.KengIndexHandler)

	return s
}
