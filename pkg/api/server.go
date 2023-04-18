package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/go-gin-google-oauth2/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(authHandler *handler.AuthHandler) *ServerHTTP {

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("./views/*.html")
	// setup routes
	login := engine.Group("/login")
	{
		login.GET("/", authHandler.LoginPage)                           // for render the login page
		login.GET("/auth/", authHandler.StartAuthentication)            // for start the google's login
		login.GET("/auth/google/callback", authHandler.CallbackHandler) // call back api for google to send user details to backend
	}

	return &ServerHTTP{
		engine: engine,
	}
}

func (c *ServerHTTP) Start() {
	c.engine.Run(":8000")
}
