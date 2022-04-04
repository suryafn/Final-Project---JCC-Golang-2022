package controllers

import (
	"net/http"
	"time"

	"FinalProject-JCCGolang2022/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productInput struct {
	Name              string `json:"title"`
	ExpiredYear       int    `json:"year"`
	ProductCategoryID uint   `json:"product_category_id"`
}

// GetAllProduct godoc
// @Summary Get all products.
// @Description Get a list of Products.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Product
// @Router /products [get]
func GetAllMovie(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateMovie godoc
// @Summary Create New Product.
// @Description Creating a new Product.
// @Tags Product
// @Param Body body productInput true "the body to create a new product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Product
// @Router /products [post]
func CreateMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input productInput
	var rating models.ProductCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ProductCategoryID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ageRatingCategoryID not found!"})
		return
	}

	// Create Product
	product := models.Product{Title: input.Name, Year: input.ExpiredYear, ProductCategoryID: input.ProductCategoryID}
	db.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetMovieById godoc
// @Summary Get Product.
// @Description Get a Product by id.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func GetMovieById(c *gin.Context) { // Get model if exist
	var product models.Product

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateMovie godoc
// @Summary Update Product.
// @Description Update product by id.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "product id"
// @Param Body body productInput true "the body to update an product"
// @Success 200 {object} models.Product
// @Router /products/{id} [patch]
func UpdateMovie(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product models.Product
	var rating models.ProductCategory
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ProductCategoryID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ageRatingCategoryID not found!"})
		return
	}

	var updatedInput models.Product
	updatedInput.Title = input.Name
	updatedInput.Year = input.ExpiredYear
	updatedInput.ProductCategoryID = input.ProductCategoryID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&product).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteMovie godoc
// @Summary Delete one product.
// @Description Delete a product by id.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} map[string]boolean
// @Router /product/{id} [delete]
func DeleteMovie(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
