package config

import (
	"log"
	"github.com/cloudinary/cloudinary-go/v2"
	
)

var CLD *cloudinary.Cloudinary
func ConncetCloudinary(){
	var err error
	CLD, err = cloudinary.NewFromParams("eloy411", "655444365249817", "fRaR-rNpRGZjkCZed9Yl81YfcKk")

	if err != nil {
		log.Fatal(err)
	}

}



