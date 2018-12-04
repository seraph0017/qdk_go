package service

import (
	"fmt"
	"github.com/bluebreezecf/opentsdb-goclient/client"
	"github.com/bluebreezecf/opentsdb-goclient/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math"
	"strconv"
	"time"
)

type (
	KengMgr interface {
		GetAllKeng() ([]KengResult, error)
	}

	KengConfig struct {
		MysqlUri string
		TsdbUri  string
	}

	KengService struct {
		Base
	}

	KengResult struct {
		Id          int    `json:"id"`
		DeviceId    string `json:"device_id"`
		Comment     string `json:"comment"`
		Gender      int    `json:"gender"`
		Floor       int    `json:"floor"`
		Location    string `json:"location"`
		IsAvailable bool   `json:"is_available"`
	}
)

const (
	MAIN_TABLE   = "keng"
	SELECT_ITEMS = `
		keng.id as id,
		pi.device_id as device_id,
		keng.comment as comment,
        keng.gender as gender,
		keng.floor as floor,
        keng.location as location
	`
	JOIN_TABLE = "left join pi on keng.id = pi.keng_id"

	NO_BODY_LINE = 124.00
	RRANGE       = 20.00
)

func (k *KengService) GetAllKeng() ([]KengResult, error) {
	res := k.Conn.
		Table(MAIN_TABLE).
		Select(SELECT_ITEMS).
		Joins(JOIN_TABLE)
	kr := make([]KengResult, 0)
	res.Scan(&kr)
	tags := make(map[string]string)
	for index, lkr := range kr {
		tags["device_id"] = lkr.DeviceId
		isAvailable := k.isAvailable(tags)
		kr[index].IsAvailable = isAvailable
	}
	return kr, nil

}

func (k *KengService) isAvailable(tags map[string]string) bool {
	queryLast := client.QueryLastParam{}
	sbq := make([]client.SubQueryLast, 0)
	sbqq := client.SubQueryLast{
		Metric: "distance",
		Tags:   tags,
	}
	sbq = append(sbq, sbqq)
	queryLast.Queries = sbq

	if queryResp, err := k.TsdbCli.QueryLast(queryLast); err != nil {
		fmt.Printf("Error occurs when querying: %v", err)
	} else {
		for _, cnt := range queryResp.QueryRespCnts {
			fmt.Printf("%+v\n", cnt.Value)
			distance, err := strconv.ParseFloat(cnt.Value, 64)
			if err != nil {
				fmt.Printf("err => %s\n", err)
			}
			if math.Abs(distance-NO_BODY_LINE) < RRANGE {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func (k *KengService) getDuration(tags map[string]string) {
	endTime := time.Now().Unix()
	startTime := time.Now().Add(-(7 * 24 * 60 * 60 * time.Second)).Unix()
	queryParam := client.QueryParam{
		Start: startTime,
		End:   endTime,
	}
	subqueries := make([]client.SubQuery, 0)
	subQuery := client.SubQuery{
		Aggregator: "none",
		Metric:     "distance",
		Tags:       tags,
	}
	subqueries = append(subqueries, subQuery)
	queryParam.Queries = subqueries

	if queryResp, err := k.TsdbCli.Query(queryParam); err != nil {
		fmt.Printf("Error occurs when querying: %v", err)
	} else {
		fmt.Printf("%s\n", queryResp.String())
		for _, cnt := range queryResp.QueryRespCnts {
			for k, v := range cnt.Dps {
				fmt.Printf("%s : %s\n", k, v)
			}
		}
	}
}

func NewKengMgr(cfg *KengConfig) (KengMgr, error) {
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
	return &KengService{Base: Base{TsdbCli: TsdbCli, Conn: MysqlConn}}, nil
}
