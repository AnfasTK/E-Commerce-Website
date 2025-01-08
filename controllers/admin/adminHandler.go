package controllers

import (
	"fmt"
	"net/http"

	"github.com/anfastk/E-Commerce-Website/config"
	"github.com/anfastk/E-Commerce-Website/middleware"
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
			"status": "Fail",
			"error":  err.Error(),
			"code":   400,
		})
		return
	}

	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Fail",
			"error":  "Email and Passowerd are required",
			"code":   400,
		})
		return
	}

	if err := config.DB.Where("email = ?", input.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Fail",
			"error":  "invalid Email",
			"code":   400,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Fail",
			"error":  "Invalid Passowrd",
			"code":   401,
		})
		return
	}

	token, err := middleware.GenerateJWT(admin.ID, admin.Email, admin.Password)
	if err != nil {
		fmt.Println("Error for generating JWT tokens")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  "Failed to generate JWT tokens",
			"code":   "500",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Login Successful",
		"token":  token,
		"code":   "200",
	})
}
