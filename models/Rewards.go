package models

type RewardsShop struct {
	IdRewardShop int `gorm:"primaryKey"`
	Name         string
	Price        int
	Url          string
}

type RewardsUsers struct {
	IdData       int `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Id_User      string
	Name         string
	ParamX       float64
	ParamY       float64
	ParamaZ      float64
	IdRewardShop string
	Url          string
}

type Coins struct {
	IdUser string
	Coins  int
}
