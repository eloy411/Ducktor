package models

type RewardsShop struct {
	IdRewardShop string `gorm:"primaryKey"`
	Name         string
	Price        int
}

type RewardsUsers struct {
	Id           string `gorm:"primaryKey"`
	Id_User      string
	Name         string
	IdRewardShop string
}



