package controller

import (
	"fmt"
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

func (p *ProductController) UpdateInElasticSearch(ctx *gin.Context) {
	var updateFields map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateFields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	docID := ctx.Param("id")
	err := p.ProductService.UpdateInElasticSearch(docID, updateFields)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func (p *ProductController) CreateInElasticSearch(ctx *gin.Context) {
	var document map[string]interface{}
	if err := ctx.ShouldBindJSON(&document); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(document)
	docID, err := p.ProductService.CreateInElasticSearch(&document)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": docID})
}

func (p *ProductController) GetAllInElasticSearch(ctx *gin.Context) {
	documents, err := p.ProductService.GetAllInElasticSearch()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, documents)
}

func (p *ProductController) GetByIdInElasticSearch(ctx *gin.Context) {
	docID := ctx.Param("id")
	document, err := p.ProductService.GetByIdInElasticSearch(docID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, document)
}

func (p *ProductController) DeleteInElasticSearch(ctx *gin.Context) {
	docID := ctx.Param("id")
	err := p.ProductService.DeleteInElasticSearch(docID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
