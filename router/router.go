package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rama4zis/go-mygram-aplication/controllers/commentcontroller"
	"github.com/rama4zis/go-mygram-aplication/controllers/usercontroller"
	"github.com/rama4zis/go-mygram-aplication/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", usercontroller.Register)
		userRouter.POST("/login", usercontroller.Login)

	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", commentcontroller.CreateComment)
	}

	return r
}
