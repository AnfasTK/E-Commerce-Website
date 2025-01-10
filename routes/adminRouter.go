package routes

import (
	controllers "github.com/anfastk/E-Commerce-Website/controllers/admin"
	"github.com/anfastk/E-Commerce-Website/middleware"
	"github.com/gin-gonic/gin"
)

var RoleAdmin = "Admin"

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.GET("/login", controllers.ShowLoginPage)
		admin.POST("/login", controllers.AdminLoginHandler)
	}
	userRoutes := r.Group("/admin/users")
	{
		userRoutes.GET("/",middleware.AuthMiddleware(RoleAdmin), controllers.ListUsers)
		userRoutes.POST("/:id/block",middleware.AuthMiddleware(RoleAdmin), controllers.BlockUser)
		userRoutes.POST("/:id/delete",middleware.AuthMiddleware(RoleAdmin),controllers.DeleteUser)
	}
	category:=r.Group("/admin/category")
	{
		category.GET("/",middleware.AuthMiddleware(RoleAdmin),controllers.ListCategory)
		category.POST("/:id/edit",middleware.AuthMiddleware(RoleAdmin),controllers.EditCategory)
		category.POST("/add",middleware.AuthMiddleware(RoleAdmin),controllers.AddCategory)
		category.POST("/:id/delete",middleware.AuthMiddleware(RoleAdmin),controllers.DeleteCategory)
	}
}
