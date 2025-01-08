package controllers

import (
	"net/http"

	"github.com/anfastk/E-Commerce-Website/config"
	models "github.com/anfastk/E-Commerce-Website/models/adminModels"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AdminLoginHandler(c *gin.Context) {
	var admin models.AdminModel
	var input AdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := config.DB.Where("email = ?", input.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid Email",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Passowrd",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Login Successful",
	})
}
