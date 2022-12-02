package models

type StatesDucktorDaily struct {
	Id           int `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	IdUser string
	Date   string
	good   int
	average int
	danger int
}