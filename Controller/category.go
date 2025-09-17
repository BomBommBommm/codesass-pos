package controller

import (
	"codezard-pos/db"
	"codezard-pos/dto"
	"codezard-pos/model"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct{}

func (c Category) FindAll(ctx *gin.Context) {
	var categories []model.Category
	db.Conn.Find(&categories)

	var result []dto.CategoryResponse
	for _,category := range categories {
		result = append(result, dto.CategoryResponse{
			ID: category.ID,
			Name: category.Name,
		})
	}
	ctx.JSON(http.StatusOK, result)
}

func (c Category) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var category model.Category
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID: category.ID,
		Name: category.Name,
	})
}
//POST /category <<-{"name":"Something"} (JSON)
func (c Category) Create(ctx *gin.Context) {

	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := model.Category{
		Name: form.Name,
	}

	if err := db.Conn.Create(&category).Error; err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ID, Name
	ctx.JSON(http.StatusCreated, dto.CategoryResponse{
		ID: category.ID,
		Name: category.Name,
	})
}

func (c Category) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category model.Category
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	category.Name = form.Name
	db.Conn.Save(&category)
	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID: category.ID,
		Name: category.Name,
	})
}

func (c Category) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	db.Conn.Delete(&model.Category{}, id)					//SOFT DELETE
	// db.Conn.Unscoped().Delete(&model.Category{}, id)		//DELETE

	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}