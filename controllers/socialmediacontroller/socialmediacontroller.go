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
	var socialMedia models.SocialMedia
	db := models.DB
	socialMediaId := c.Param("socialMediaId")
	// turn string to uint model
	var socialMediaIdUint uint
	socialMediaIdUint = helpers.StringToUint(socialMediaId)

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	if contentType == "application/json" {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	socialMedia.Id = socialMediaIdUint

	c.JSON(http.StatusOK, gin.H{"data": socialMedia})
	return

	// set user id
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	socialMedia.UserId = userId

	// Create SocialMedia
	err := db.Debug().Save(&socialMedia).Error

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
