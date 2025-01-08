package main

import (
	"github.com/anfastk/E-Commerce-Website/config"
	"github.com/anfastk/E-Commerce-Website/routes"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.Static("static","./static")
	config.DBconnect()
	config.LoadEnvFile()
	config.SyncDatabase()
}

func main(){
	routes.AdminRoutes(r)
	r.Run()
}