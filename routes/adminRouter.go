package routes

import (
	controllers "github.com/anfastk/E-Commerce-Website/controllers/admin"
	"github.com/anfastk/E-Commerce-Website/middleware"
	"github.com/gin-gonic/gin"
)

var RoleAdmin = "Admin"

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.Use(middleware.NoCacheMiddleware())
	{
		admin.GET("/login", controllers.ShowLoginPage)
		admin.POST("/login", controllers.AdminLoginHandler)
	}
	product := r.Group("/admin/products")
	product.Use(middleware.AuthMiddleware(RoleAdmin))
	product.Use(middleware.NoCacheMiddleware())
	{
		product.GET("/", controllers.ShowProductsAdmin)
		product.GET("/main/add", controllers.ShowAddMainProduct)
		product.POST("main/add", controllers.AddMainProductDetails)
	}
	adminUser := r.Group("/admin/users")
	adminUser.Use(middleware.AuthMiddleware(RoleAdmin))
	adminUser.Use(middleware.NoCacheMiddleware())
	{
		adminUser.GET("/", controllers.ListUsers)
		adminUser.POST("/:id/block", controllers.BlockUser)
		adminUser.POST("/:id/delete", controllers.DeleteUser)
	}
	category := r.Group("/admin/category")
	category.Use(middleware.AuthMiddleware(RoleAdmin))
	category.Use(middleware.NoCacheMiddleware())
	{
		category.GET("/", controllers.ListCategory)
		category.POST("/:id/edit", controllers.EditCategory)
		category.POST("/add", controllers.AddCategory)
		category.POST("/:id/delete", controllers.DeleteCategory)
	}
}
