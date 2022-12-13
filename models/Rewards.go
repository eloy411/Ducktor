package models

type RewardsShop struct {
	IdRewardShop int `gorm:"primaryKey"`
	Name         string
	ParamX       float64
	ParamY       float64
	ParamaZ      float64
	RotationX    float64 `gorm:"default:0.0"`
	RotationY    float64 `gorm:"default:0.0"`
	Price        int
	Url          string
	UrlImagen	 string
	UrlIntegrado string `gorm:"default:''"`
}

type RewardsUsers struct {
	IdData       int `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Id_User      string
	Name         string
	ParamX       float64
	ParamY       float64
	ParamaZ      float64
	RotationX    float64 `gorm:"default:0.0"`
	RotationY    float64 `gorm:"default:0.0"`
	IdRewardShop int 
	Url          string
	UrlIntegrado string `gorm:"default:''"`
}

type Coins struct {
	IdUser string 
	Coins  int
}
