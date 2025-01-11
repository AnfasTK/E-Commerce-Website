package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowProductsAdmin(c *gin.Context) {
	c.HTML(http.StatusOK, "productPage.html", gin.H{})
}

func ShowAddProductPage(c *gin.Context){
	c.HTML(http.StatusOK, "addNewMainProductDetails.html", gin.H{})
}