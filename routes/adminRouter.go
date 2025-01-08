package routes

import (
	controllers "github.com/anfastk/E-Commerce-Website/controllers/admin"
	"github.com/gin-gonic/gin"
)

var RoleAdmin = "Admin"

func AdminRoutes(r *gin.Engine){
	
	r.LoadHTMLGlob("views/Admin/*")

	admin:= r.Group("/admin")
	{
		admin.GET("/login",controllers.ShowLoginPage)
		admin.POST("/login",controllers.AdminLoginHandler)
	}
}
