package photocontroller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rama4zis/go-mygram-aplication/helpers"
	"github.com/rama4zis/go-mygram-aplication/models"
)

func CreatePhoto(c *gin.Context) {
	db := models.DB

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create Photo!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.Id,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserId,
		"created_at": Photo.CreatedAt,
	})
}

func UpdatePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	db := models.DB

	Photo := models.Photo{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := db.Debug().Where("id = ?", photoID).First(&Photo).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to find Photo!",
			"error":   err.Error(),
		})
		return
	}

	if Photo.UserId != userID {
		c.JSON(401, gin.H{
			"message": "Unauthorized!",
		})
		return
	}

	contentType := helpers.GetContentType(c)

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err = db.Debug().Save(&Photo).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to update Photo!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.Id,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserId,
		"updated_at": Photo.UpdatedAt,
	})

	// c.JSON(http.StatusOK, Photo)
}

func GetAllPhotos(c *gin.Context) {
	// db := models.DB
	contentType := helpers.GetContentType(c)

	photo := []models.Photo{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	models.DB.Preload("UserPhoto").Find(&photo)

	c.JSON(http.StatusOK, photo)
}

func DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	db := models.DB

	Photo := models.Photo{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := db.Debug().Where("id = ?", photoID).First(&Photo).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to find Photo!",
			"error":   err.Error(),
		})
		return
	}

	if Photo.UserId != userID {
		c.JSON(401, gin.H{
			"message": "Unauthorized!",
		})
		return
	}

	err = db.Debug().Delete(&Photo).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to delete Photo!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo deleted successfully!",
	})
}
