package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"os"
	"qdk/qdk.v1/model"
	"time"
)

type Manager struct {
	Db *gorm.DB
}

func (m *Manager) createDb(c *cli.Context) error {
	m.Db.CreateTable(model.Pi{})
	m.Db.CreateTable(model.Keng{})
	return nil
}

func (m *Manager) clearDb(c *cli.Context) error {
	m.Db.DropTable(model.Pi{})
	m.Db.DropTable(model.Keng{})
	return nil
}

func (m *Manager) initDb(c *cli.Context) error {
	m.clearDb(c)
	m.createDb(c)
	k1 := model.Keng{
		Comment:  "12楼男厕所",
		Gender:   1,
		Floor:    12,
		Location: "左边",
		Base: model.Base{
			CreationTime: time.Now(),
			ModifiedTime: time.Now(),
		},
	}
	k2 := model.Keng{
		Comment:  "12楼男厕所",
		Gender:   1,
		Floor:    12,
		Location: "右边",
		Base: model.Base{
			CreationTime: time.Now(),
			ModifiedTime: time.Now(),
		},
	}
	m.Db.Create(&k1)
	m.Db.Create(&k2)
	p1 := model.Pi{
		DeviceId: "fujun-1",
		Ip:       "0.0.0.0",
		MacId:    "11111",
		KengId:   k1.ID,
		Comment:  "12楼左边男厕所的树莓派",
		Alias:    "坑001",
		Base: model.Base{
			CreationTime: time.Now(),
			ModifiedTime: time.Now(),
		},
	}
	p2 := model.Pi{
		DeviceId: "fujun-2",
		Ip:       "0.0.0.0",
		MacId:    "11111",
		KengId:   k2.ID,
		Comment:  "12楼左边男厕所的树莓派",
		Alias:    "坑002",
		Base: model.Base{
			CreationTime: time.Now(),
			ModifiedTime: time.Now(),
		},
	}
	m.Db.Create(&p1)
	m.Db.Create(&p2)
	return nil
}

func main() {
	viper.SetConfigName("dev")
	viper.AddConfigPath("./qdk.v1/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", viper.GetString("mysql_uri"))
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
	m := &Manager{db}
	app := cli.NewApp()
	app.Name = "qdk"
	app.Usage = "让所有人都可以监控厕所"

	app.Commands = []cli.Command{
		{
			Name:   "createdb",
			Action: m.createDb,
		},
		{
			Name:   "cleardb",
			Action: m.clearDb,
		},
		{
			Name:   "initdb",
			Action: m.initDb,
		},
	}
	app.Run(os.Args)
}
