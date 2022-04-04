package controllers

import (
	"net/http"
	"time"

	"FinalProject-JCCGolang2022/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetAllAgeRatingCategory godoc
// @Summary Get all ProductCategory.
// @Description Get a list of ProductCategory.
// @Tags ProductCategory
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.ProductCategory
// @Router /product-categories [get]
func GetAllRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var ratings []models.ProductCategory
	db.Find(&ratings)

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// CreateAgeRatingCategory godoc
// @Summary Create New ProductCategory.
// @Description Creating a new ProductCategory.
// @Tags ProductCategory
// @Param Body body productCategoryInput true "the body to create a new ProductCategory"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.ProductCategory
// @Router /product-categories [post]
func CreateRating(c *gin.Context) {
	// Validate input
	var input productCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Product
	product := models.ProductCategory{Name: input.Name, Description: input.Description}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetAgeRatingCategoryById godoc
// @Summary Get ProductCategory.
// @Description Get an ProductCategory by id.
// @Tags ProductCategory
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "ProductCategory id"
// @Success 200 {object} models.ProductCategory
// @Router /product-categories/{id} [get]
func GetRatingById(c *gin.Context) { // Get model if exist
	var product models.ProductCategory

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetMovieByAgeRatingCategoryId godoc
// @Summary Get Movies.
// @Description Get all Movies by AgeRatingCategoryId.
// @Tags ProductCategory
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "ProductCategory id"
// @Success 200 {object} []models.Product
// @Router /product-categories/{id}/movies [get]
func GetMoviesByRatingId(c *gin.Context) { // Get model if exist
	var movies []models.Product

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("age_rating_category_id = ?", c.Param("id")).Find(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// UpdateAgeRatingCategory godoc
// @Summary Update ProductCategory.
// @Description Update ProductCategory by id.
// @Tags ProductCategory
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "ProductCategory id"
// @Param Body body productCategoryInput true "the body to update age product category"
// @Success 200 {object} models.ProductCategory
// @Router /product-categories/{id} [patch]
func UpdateRating(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product models.ProductCategory
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input productCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.ProductCategory
	updatedInput.Name = input.Name
	updatedInput.Description = input.Description
	updatedInput.UpdatedAt = time.Now()

	db.Model(&product).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteAgeRatingCategory godoc
// @Summary Delete one ProductCategory.
// @Description Delete a ProductCategory by id.
// @Tags ProductCategory
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "ProductCategory id"
// @Success 200 {object} map[string]boolean
// @Router /product-categories/{id} [delete]
func DeleteRating(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.ProductCategory
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
