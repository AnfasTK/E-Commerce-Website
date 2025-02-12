package helper

import (
	"github.com/anfastk/E-Commerce-Website/config"
	"github.com/anfastk/E-Commerce-Website/models"
)

type productsResponse struct {
	ID             uint     `json:"id"`
	ProductName    string   `json:"product_name"`
	ProductSummary string   `json:"product_summary"`
	SalePrice      float64  `json:"sale_price"`
	RegularPrice   float64  `json:"regular_price"`
	Images         []string `json:"images"`
}

func RelatedProducts(categoryID uint) ([]productsResponse, error) {
	var products []models.ProductVariantDetails
	if err := config.DB.Preload("VariantsImages", "is_deleted = ?", false).
		Where("category_id = ? AND is_deleted = ?", categoryID, false).
		Limit(20).
		Find(&products).Error; err != nil {
		return nil, err
	}

	var response []productsResponse
	for _, product := range products {
		var images []string
		for _, image := range product.VariantsImages {
			images = append(images, image.ProductVariantsImages)
		}
		response = append(response, productsResponse{
			ID:             product.ID,
			ProductName:    product.ProductName,
			ProductSummary: product.ProductSummary,
			SalePrice:      product.SalePrice,
			RegularPrice:   product.RegularPrice,
			Images:         images,
		})
	}
	return response, nil
}