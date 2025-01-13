package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/anfastk/E-Commerce-Website/config"
	"github.com/anfastk/E-Commerce-Website/models"
	"github.com/anfastk/E-Commerce-Website/utils"
	"github.com/gin-gonic/gin"
)

func ShowProductsAdmin(c *gin.Context) {
	c.HTML(http.StatusOK, "productPage.html", gin.H{})
}

func ShowAddMainProduct(c *gin.Context) {
	var categories []models.Categories
	if err := config.DB.Where("is_deleted = ? AND status = ?", false, "Active").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Failed to fetch categories",
			"error": err.Error(),
			"code":    500,
		})
		return
	}

	c.HTML(http.StatusOK, "addNewMainProductDetails.html", gin.H{
		"categories": categories,
	})
}
func AddMainProductDetails(c *gin.Context) {
	tx := config.DB.Begin()
	categoryID, err := strconv.ParseInt(c.PostForm("category"), 10, 64)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Invalid category ID",
			"error": err.Error(),
			"code":    500,
		})
		return
	}
	product := models.ProductDetail{
		ProductName:    c.PostForm("product_name"),
		CategoryID:     uint(categoryID),
		BrandName:      c.PostForm("brand_name"),
		IsCODAvailable: c.PostForm("cod_available") == "YES",
		IsReturnable:   c.PostForm("return_available") == "YES",
	}
	if err:=tx.Create(&product).Error;err!=nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Failed to save product",
			"error": err.Error(),
			"code":    500,
		})
		return
	}
	cld:=config.InitCloudinary()

	form,_:=c.MultipartForm()
	fmt.Println("====================",form.File["product_image"],"========================")
	if form !=nil{
		if productImage,ok:=form.File["product_image"];ok && len(productImage)>0 {
			fileHeader:=productImage[0]
			file,_:=fileHeader.Open()
			defer file.Close()

			url,err:=utils.UploadImageToCloudinary(file,fileHeader,cld)
			if err !=nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Internal Server Error",
					"message": "Failed to upload product image",
					"error": err.Error(),
					"code":    500,
				})
				return
			}
			image:=models.ProductImage{
				ProductImages: url,
				ProductID: product.ID,
			}
			fmt.Println(image)
			if err:=tx.Create(&image).Error;err!=nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Internal Server Error",
					"message": "Failed to upload product image",
					"error": err.Error(),
					"code":    500,
				})
				return
			}
		}
	}
	tx.Commit()
	c.Redirect(http.StatusFound,"/admin/products/main/details")
}
