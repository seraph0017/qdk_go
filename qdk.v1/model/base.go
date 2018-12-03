package model

import (
	"time"
)

type (
	Base struct {
		ID           int       `gorm:"primary_key" json:"id"`
		CreationTime time.Time `json:"creation_time"`
		ModifiedTime time.Time `json:"modified_time"`
		Ext          string    `json:"ext"`
	}
)
