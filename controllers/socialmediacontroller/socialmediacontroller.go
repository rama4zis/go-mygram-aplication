package socialmediacontroller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rama4zis/go-mygram-aplication/helpers"
	"github.com/rama4zis/go-mygram-aplication/models"
)

func CreateSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia
	db := models.DB

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	if contentType == "application/json" {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	// c.JSON(http.StatusOK, gin.H{"data": socialMedia})
	// return

	// set user id
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	socialMedia.UserId = userId

	// Create SocialMedia
	err := db.Debug().Create(&socialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	returnData := map[string]interface{}{
		"id":               socialMedia.Id,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserId,
		"craeted_at":       socialMedia.CreatedAt,
	}

	c.JSON(http.StatusCreated, returnData)
}

func GetAllSocialMedias(c *gin.Context) {
	var socialMedias []models.SocialMedia

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	result := models.DB.Where("user_id = ?", userId).Find(&socialMedias)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Failed to fetch social media!",
		})
		return
	}

	userData = c.MustGet("userData").(jwt.MapClaims)
	userId = uint(userData["id"].(float64))
	thisUser := models.User{}
	models.DB.Where("id = ?", userId).First(&thisUser)

	var finalData []map[string]interface{}
	for i := range socialMedias {
		returnData := map[string]interface{}{
			"id":               socialMedias[i].Id,
			"name":             socialMedias[i].Name,
			"social_media_url": socialMedias[i].SocialMediaUrl,
			"UserId":           socialMedias[i].UserId,
			"craetedAt":        socialMedias[i].CreatedAt,
			"updatedAt":        socialMedias[i].UpdatedAt,
			"User": struct {
				Id              uint   `json:"id"`
				Username        string `json:"username"`
				ProfileImageURL string `json:"profile_image_url"`
			}{
				Id:              thisUser.Id,
				Username:        thisUser.Username,
				ProfileImageURL: "Ini url dari mana?",
			},
		}
		finalData = append(finalData, returnData)
	}

	c.JSON(200, gin.H{
		"social_meidas": finalData,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	socialMediaId := c.Param("socialMediaId")
	db := models.DB
	var SocialMedia models.SocialMedia

	// check if social media exist
	err := db.Debug().First(&SocialMedia, socialMediaId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Social media not found!"})
		return
	}

	// check if user is owner of social media
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	if SocialMedia.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		return
	}

	// update social media
	contentType := helpers.GetContentType(c)
	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err = db.Debug().Save(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	returnData := map[string]interface{}{
		"id":               SocialMedia.Id,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserId,
		"craeted_at":       SocialMedia.CreatedAt,
	}

	c.JSON(http.StatusCreated, returnData)
}

func DeleteSocialMedia(c *gin.Context) {
	socialMediaId := c.Param("socialMediaId")
	db := models.DB
	var SocialMedia models.SocialMedia

	// check if social media exist
	err := db.Debug().First(&SocialMedia, socialMediaId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Social media not found!"})
		return
	}

	// check if user is owner of social media
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	if SocialMedia.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		return
	}

	// delete social media
	err = db.Debug().Delete(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your social media has been successfully deleted!"})
}
