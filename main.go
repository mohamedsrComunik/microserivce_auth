package main

import (
	"github.com/mohamedsrComunik/microserivce_auth/database"
	"github.com/mohamedsrComunik/microserivce_auth/handler"
	"github.com/mohamedsrComunik/microserivce_auth/jwt"
	"github.com/mohamedsrComunik/microserivce_auth/repository"
	"github.com/mohamedsrComunik/microserivce_auth/router"
)

func main() {
	db := database.Connexion()
	defer db.Close()

	repo := repository.New(db)
	jwt := jwt.TokenService{}
	var repoInterface repository.Repository
	repoInterface = &repo
	var handlerInterface handler.Handler
	handler := handler.New(repoInterface, &jwt)
	handlerInterface = &handler

	//Init router
	routes := router.New(handlerInterface)
	r := routes.SetRoutes()
	r.Run()
}
