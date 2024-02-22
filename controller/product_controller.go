package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhvdt/vcs-week2/service/product"
	"net/http"
)

type ProductController struct {
	ProductService product.ProductService
}

func NewProductController(productService product.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (p *ProductController) GetProductsWithSupplier(ctx *gin.Context) {
	products, err := p.ProductService.GetProductsWithSupplier()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProductAboveAveragePrice(ctx *gin.Context) {
	products, err := p.ProductService.GetProductAboveAveragePrice()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
