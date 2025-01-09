package config

import (
	"log"

	models "github.com/anfastk/E-Commerce-Website/models"
)

func SyncDatabase() {
	err:=DB.AutoMigrate(&models.AdminModel{})
	if err != nil{
		log.Fatalf("Failed to migrate models: %v",err)
	}
	log.Println("Models migrated")
}