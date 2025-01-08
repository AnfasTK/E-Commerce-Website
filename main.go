package main

import (
	"github.com/anfastk/E-Commerce-Website/config"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("static","./static")
	config.DBconnect()
	config.LoadEnvFile()
	config.SyncDatabase()
}

func main(){
	r.Run()
}