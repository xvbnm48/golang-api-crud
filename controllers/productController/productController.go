package productController

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/golang-api-crud/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var product []models.Product

	models.DB.Find(&product)
	c.JSON(200, gin.H{
		"data": product,
	})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(404, gin.H{
				"message": "Record not found",
			})
			return
		default:
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
		}
	}
	c.JSON(200, gin.H{
		"data": product,
	})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	models.DB.Create(&product)
	c.JSON(200, gin.H{
		"data": product,
	})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "Record not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Record updated successfully",
	})
}

func Delete(c *gin.Context) {
	var product models.Product
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "Tidak dapat menghapus record",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Record deleted successfully",
	})
}
