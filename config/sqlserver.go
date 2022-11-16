package config

import (
	"fmt"
	"log"



	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DSN = "sqlserver://:@127.0.0.1:1433?database=tienda"
var DB *gorm.DB

func ConnectDB() {

	var err error
	
	DB, err = gorm.Open(sqlserver.Open(DSN), &gorm.Config{})

	if err != nil {
		fmt.Println("no se ha podido conectar")
	}else{
		log.Println("DB connected")
	}
}
