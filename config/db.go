package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect(){
	var err error
	dsn:="host=localhost user=postgres password=131020 dbname=laptix_ecommerce_website port=5432"
	DB,err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatalf("Failed to connect database: %v",err)
	}
	log.Println("Database connected successfully")
}