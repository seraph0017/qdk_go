package model

type Pi struct {
	Base
	DeviceId   string  `json:"device_id"`
	Ip         string  `json:"ip"`
	MacId      string  `json:"mac_id"`
	KengId     int     `json:"keng_id"`
	Comment    string  `json:"comment"`
	Alias      string  `json:"alias"`
	Status     int     `json:"-" gorm:"default:0"`
	NobodyLine float64 `json:"nobody_line" gorm:"default:124.00"`
	RRange     float64 `json:"rrange" gorm:"default:20.00"`
}
