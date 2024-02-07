package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/thanhvdt/vcs-week2/model"
	"github.com/thanhvdt/vcs-week2/service"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	log.Info().Msg("Creating Category")
	err := ctx.ShouldBindJSON(&model.Category{})
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	category, er := c.categoryService.Create(&model.Category{})
	if er != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := gin.H{
		"code":   200,
		"status": "Created",
		"data":   category,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

func (c *CategoryController) ReadAllCategory(ctx *gin.Context) {
	log.Info().Msg("Reading All Category")
	categories, err := c.categoryService.ReadAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := gin.H{
		"code":   200,
		"status": "OK",
		"data":   categories,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

func (c *CategoryController) ReadCategoryByID(ctx *gin.Context) {
	log.Info().Msg("Reading Category By ID")
	categoryID := ctx.Param("categoryID")
	category, err := c.categoryService.ReadByID(categoryID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := gin.H{
		"code":   200,
		"status": "OK",
		"data":   category,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	log.Info().Msg("Updating Category")
	err := ctx.ShouldBindJSON(&model.Category{})
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	category, er := c.categoryService.Update(&model.Category{})
	if er != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := gin.H{
		"code":   200,
		"status": "Updated",
		"data":   category,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	log.Info().Msg("Deleting Category")
	categoryID := ctx.Param("categoryID")
	err := c.categoryService.Delete(categoryID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := gin.H{
		"code":   200,
		"status": "Deleted",
		"data":   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}
