package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rama4zis/go-mygram-aplication/controllers/commentcontroller"
	"github.com/rama4zis/go-mygram-aplication/controllers/socialmediacontroller"

	"github.com/rama4zis/go-mygram-aplication/controllers/photocontroller"
	"github.com/rama4zis/go-mygram-aplication/controllers/usercontroller"
	"github.com/rama4zis/go-mygram-aplication/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	usersRouter := r.Group("/users")
	{
		usersRouter.POST("/register", usercontroller.Register)
		usersRouter.POST("/login", usercontroller.Login)
		usersRouter.PUT("/", middlewares.Authentication(), usercontroller.Update)
		usersRouter.DELETE("/", middlewares.Authentication(), usercontroller.Delete)

	}

	photosRouter := r.Group("/photos")
	{
		photosRouter.Use(middlewares.Authentication())
		photosRouter.POST("/", middlewares.Authentication(), photocontroller.CreatePhoto)
		photosRouter.PUT(("/:photoId"), middlewares.Authentication(), photocontroller.UpdatePhoto)
		photosRouter.DELETE(("/:photoId"), middlewares.Authentication(), photocontroller.DeletePhoto)

		photosRouter.GET("/", photocontroller.GetAllPhotos)
	}

	commentsRouter := r.Group("/comments")
	{
		commentsRouter.Use(middlewares.Authentication())
		commentsRouter.POST("/", commentcontroller.CreateComment)
		commentsRouter.PUT(("/:commentId"), middlewares.Authentication(), commentcontroller.UpdateComment)
		commentsRouter.DELETE(("/:commentId"), middlewares.Authentication(), commentcontroller.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.POST("/", middlewares.Authentication(), socialmediacontroller.CreateSocialMedia)
		socialMediaRouter.PUT(("/:socialMediaId"), middlewares.Authentication(), socialmediacontroller.UpdateSocialMedia)
		socialMediaRouter.DELETE(("/:socialMediaId"), middlewares.Authentication(), socialmediacontroller.DeleteSocialMedia)

		socialMediaRouter.GET("/", middlewares.Authentication(), socialmediacontroller.GetAllSocialMedias)
	}

	return r
}
