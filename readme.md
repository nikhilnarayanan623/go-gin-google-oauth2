# `Google Authentication` Implimentaion Using `Go programing` with `gin` framework and `goth google` package

# Used Packages 
1. [GIN](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
2.[GOTH]("github.com/markbates/goth") Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications.
<!-- 6. [swag](https://github.com/swaggo/swag) converts Go annotations to Swagger Documentation 2.0 with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files) -->

# To Use This `go-gin-google-oauth2` Implimentaion

# Follow these steps

1. clone my github `go-gin-google-oauth2` repository to your system
2. after cloning use follwing bash commands

``` bash commands
## these bash comands are set-up on makefilie ##
## if you did't install make then install make or use `go run ./cmd/api/`

# Step 1 :  Navigate into project directory
cd ./go-gin-google-oauth2

# Step 2 : Install needed dependencies
make deps 
#or
go mod tidy

# Step 3 : Setup Env files
    1.vist `https://console.cloud.google.com/apis/dashboard` for creating google client_id and client_secret
    #setup env
    2.GOAUTH_CLIENT_ID="Your Client Secret ID"
    3.GOAUTH_CLENT_SECRET="Your Client Secret"


# Step 4 : Run the Server
make run

# check server.go for see the api's
# open a browser and visit `localhost:8000/login`

```
