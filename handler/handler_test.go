package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohamedsrComunik/microserivce_auth/database"
	"github.com/mohamedsrComunik/microserivce_auth/handler"
	"github.com/mohamedsrComunik/microserivce_auth/jwt"
	"github.com/mohamedsrComunik/microserivce_auth/repository"
	"github.com/mohamedsrComunik/microserivce_auth/router"
	"gopkg.in/go-playground/assert.v1"
)

func TestValid(t *testing.T) {
	db := database.Connexion()

	repo := repository.New(db)
	jwt := jwt.TokenService{}
	var repoInterface repository.Repository
	repoInterface = &repo
	var handlerInterface handler.Handler
	handler := handler.New(repoInterface, &jwt)
	handlerInterface = &handler

	//Init router
	routes := router.New(handlerInterface)
	router := routes.SetRoutes()
	w := httptest.NewRecorder()
	testToken := struct {
		Token string `json:"token"`
	}{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklEIjo5LCJDcmVhdGVkQXQiOiIyMDIwLTAzLTExVDEzOjQyOjI2WiIsIlVwZGF0ZWRBdCI6IjIwMjAtMDMtMTFUMTM6NDI6MjZaIiwiRGVsZXRlZEF0IjpudWxsLCJwYXNzd29yZCI6IiQyYSQxMCRqU2p6SUVnMWg2NWtza2lRUElzVENPUloxVE4ucWlhbVB4dThnaDZKM291V2Y0aFN5MkpEYSIsImVtYWlsIjoidGVzMnRAaG90bWFpbC5mciIsImZpcnN0X25hbWUiOiIiLCJsYXN0X25hbWUiOiIiLCJSb2xlIjoiIn0sImV4cCI6MTU4NDIwMjU3OCwiaXNzIjoiYXV0aC1zcnYifQ.gvtBMtvQJKwsz51K7KnY_f4njFW89L9DUOh9gwzgSpE",
	}
	thebytes, _ := json.Marshal(testToken)
	reader := bytes.NewReader(thebytes)
	req, _ := http.NewRequest("GET", "/valid", reader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 202, w.Code)
}
