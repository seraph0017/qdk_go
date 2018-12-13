package service

import (
	"context"
	"fmt"
	"github.com/bluebreezecf/opentsdb-goclient/client"
	"github.com/bluebreezecf/opentsdb-goclient/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"qdk/qdk.v1/model"
	"time"
)

type (
	PiMgr interface {
		PostUltraSonic(ctx context.Context, input PostUltraSonicInput) error
		Validate(ctx context.Context, deviceId string) bool
	}

	PiConfig struct {
		MysqlUri string
		TsdbUri  string
		Log      *logrus.Logger
	}

	PiService struct {
		Base
	}

	PostUltraSonicInput struct {
		DeviceId string
		Distance string
	}
)

func (p *PiService) PostUltraSonic(ctx context.Context, input PostUltraSonicInput) error {
	distanceDatas := make([]client.DataPoint, 0)
	tags := make(map[string]string)
	tags["device_id"] = input.DeviceId
	data := client.DataPoint{
		Metric:    "distance",
		Timestamp: time.Now().Unix(),
		Value:     input.Distance,
	}
	data.Tags = tags

	distanceDatas = append(distanceDatas, data)
	fmt.Printf("%+v \n", distanceDatas)
	if resp, err := p.TsdbCli.Put(distanceDatas, "details"); err != nil {
		fmt.Printf("Error occurs when putting datapoints: %v", err)
	} else {
		fmt.Printf(" %s", resp.String())
	}
	return nil
}

func (p *PiService) Validate(ctx context.Context, deviceId string) bool {
	pi := model.Pi{}
	p.Conn.Where("device_id = ?", deviceId).First(&pi)
	if pi.ID != 0 {
		return true
	}
	return false
}

func NewPiMgr(cfg *PiConfig) (PiMgr, error) {
	TsdbCli, err := client.NewClient(config.OpenTSDBConfig{OpentsdbHost: cfg.TsdbUri})
	if err != nil {
		fmt.Println("err => ", err)
	}
	MysqlConn, err := gorm.Open("mysql", cfg.MysqlUri)
	if err != nil {
		fmt.Println("err => ", err)
	}
	MysqlConn.SingularTable(true)
	MysqlConn.LogMode(true)
	return &PiService{Base: Base{TsdbCli: TsdbCli, Conn: MysqlConn, Log: cfg.Log}}, nil
}
