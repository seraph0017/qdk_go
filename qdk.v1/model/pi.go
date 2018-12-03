package model

type Pi struct {
	Base
	DeviceId string `json:"device_id"`
	Ip       string `json:"ip"`
	MacId    string `json:"mac_id"`
	KengId   int    `json:"keng_id"`
	Comment  string `json:"comment"`
	Alias    string `json:"alias"`
	Status   int    `json:"-" gorm:"default:0"`
}
