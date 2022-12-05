package models

import "time"

type StatesDucktorDaily struct {
	Id      int `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	IdUser  string
	Date    time.Time `gorm:"autoCreateTime"`
	Good    int
	Average int
	Danger  int
}