package routes

import (
	controllers "github.com/anfastk/E-Commerce-Website/controllers/admin"
	"github.com/gin-gonic/gin"
)

var RoleAdmin = "Admin"

func AdminRoutes(r *gin.Engine) {

	r.LoadHTMLGlob("views/Admin/*")

	admin := r.Group("/admin")
	{
		admin.GET("/login", controllers.ShowLoginPage)
		admin.POST("/login", controllers.AdminLoginHandler)
	}
	userRoutes := r.Group("/admin/users")
	{
		userRoutes.GET("/", controllers.ListUsers)
		userRoutes.POST("/:id/block", controllers.BlockUser)
		userRoutes.POST("/:id/delete",controllers.DeleteUser)
	}
}
