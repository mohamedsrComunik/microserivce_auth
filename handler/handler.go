package handler

import (
	"log"
	"net/http"

	"github.com/mohamedsrComunik/microserivce_auth/jwt"
	"github.com/mohamedsrComunik/microserivce_auth/model"
	"github.com/mohamedsrComunik/microserivce_auth/repository"
	"gopkg.in/go-playground/validator.v9"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type handler struct {
	repo       repository.Repository
	jwtService *jwt.TokenService
}

// New to create new instance of handler
func New(repo repository.Repository, jwt *jwt.TokenService) handler {
	h := handler{repo: repo, jwtService: jwt}
	return h
}

// Handler interface
type Handler interface {
	Create(c *gin.Context)
	Auth(c *gin.Context)
	ValidateToken(c *gin.Context)
}

func (h *handler) Create(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		ValidationErrors(c, err)
		return
	}
	exists := h.repo.GetUserByEmail(&user)
	if exists == true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email exists"})
		return
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	createErr := h.repo.Create(&user)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error})
		return
	}
	token, err := h.jwtService.Encode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	// Send event to notification service
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (h *handler) Auth(c *gin.Context) {
	req := model.User{}
	if err := c.ShouldBind(&req); err != nil {
		ValidationErrors(c, err)
		return
	}
	user := model.User{
		Email: req.Email,
	}
	exists := h.repo.GetUserByEmail(&user)
	if exists != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "incorrect email"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "incorrect password"})
		return
	}
	token, err := h.jwtService.Encode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"token": token})

}

func (h *handler) ValidateToken(c *gin.Context) {
	if c.Query("token") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token not found"})
		return
	}
	// Decode token
	claims, err := h.jwtService.Decode(c.Query("token"))

	if err != nil {
		log.Println("jwt error : ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"valid": false})
		return
	}

	if claims.User.ID == 0 {
		log.Println("claims error : ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"valid": false})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"valid": true})
}

// ValidationErrors returns json response with invalid input
func ValidationErrors(c *gin.Context, err error) {
	errorsMap := gin.H{}
	for _, fieldErr := range err.(validator.ValidationErrors) {
		errorsMap[fieldErr.Field()] = fieldErr.Field() + " Invalid"
	}
	c.JSON(http.StatusInternalServerError, errorsMap)
	return
}
