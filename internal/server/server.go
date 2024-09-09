package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) RegisterRouter() http.Handler {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/", s.HelloWorldHandler)
	r.Use(middleware.ErrorHandlerMiddleware())
	r.NoRoute(middleware.HandleNotFound)

	userController := controller.NewUserController()

	r.POST("v1/auth/users/create", userController.CreateUserHttp)
	r.POST("v1/auth/users/login", userController.LoginUserHandler)
	r.POST("v1/auth/users/verify-email", userController.VerifyEmailHandler)
	r.POST("v1/auth/users/send-email", userController.SendEmailHandler)

	authorized := r.Group("/v1/auth")
	authorized.Use(middleware.VerifyToken())
	{
		authorized.POST("/users/upload-image", userController.UploadImageHandler)
	}
	r.GET("/ws", func(c *gin.Context) {
		s.ws.HandleConnections(c.Writer, c.Request)
	})

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	c.JSON(http.StatusOK, resp)

}
