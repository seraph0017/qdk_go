package model

type Keng struct {
	Base
	Comment  string `json:"comment"`
	Gender   int    `json:"gender"`
	Floor    int    `json:"floor"`
	Status   int    `json:"-" gorm:"default:0"`
	Location string `json:"location"`
}
