package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/nikhilnarayanan623/go-gin-google-oauth2/pkg/config"
)

type AuthHandler struct {
}

type response struct {
	Message   string
	Error     string
	UserID    string
	Email     string
	FirstName string
	LastName  string
}

func (c *AuthHandler) LoginPage(ctx *gin.Context) {

	ctx.HTML(200, "index.html", nil)
}

// func for start the google api authentication
func (c *AuthHandler) StartAuthentication(ctx *gin.Context) {

	// setup the google provider
	goauthClientID := config.GetCofig().GoathClientID
	gouthClientSecret := config.GetCofig().GoauthClientSecret
	callbackUrl := "http://localhost:8000/login/auth/google/callback" // copy from redirect url from google auth client creation

	// setup privider
	goth.UseProviders(
		google.New(goauthClientID, gouthClientSecret, callbackUrl, "email", "profile"),
	)

	// start the google login
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

// callback function to retrive user data's from google
func (c *AuthHandler) CallbackHandler(ctx *gin.Context) {

	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)

	if err != nil {
		ctx.HTML(500, "response.html", response{
			Message: "Faild to get user details",
			Error:   err.Error(),
		})
		return
	}
	ctx.HTML(200, "response.html", response{
		Message:   "Successfully got user details",
		Error:     "",
		UserID:    user.UserID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})

}
