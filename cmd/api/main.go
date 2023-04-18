package main

import (
	"log"

	http "github.com/nikhilnarayanan623/go-gin-google-oauth2/pkg/api"
	"github.com/nikhilnarayanan623/go-gin-google-oauth2/pkg/api/handler"
	"github.com/nikhilnarayanan623/go-gin-google-oauth2/pkg/config"
)

func main() {

	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("faild to load cofig from env")
	}

	// not initialize api because there is no many dependency injection if you want more then create it using wire
	authHandler := handler.AuthHandler{}

	server := http.NewServerHTTP(&authHandler)

	server.Start()
}
