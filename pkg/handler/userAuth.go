package handler

import (
	"go-redis-mysql_userAuth/pkg/domain"
	"go-redis-mysql_userAuth/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) *UserHandler {
	return &UserHandler{usecase}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	registerReq := domain.User{}
	if err := ctx.Bind(&registerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	user, err := h.service.CreateUser(registerReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	response := struct {
		User    domain.UserDetails
		Message string
	}{
		User:    user,
		Message: "Successfully registered",
	}
	ctx.JSON(200, response)
}

func (h *UserHandler) UserLogin(ctx *gin.Context) {
	loginReq := domain.User{}
	if err := ctx.Bind(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	user, err := h.service.UserLogin(loginReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	response := struct {
		User    domain.UserDetails
		Message string
	}{
		User:    user,
		Message: "Logged in Successfully",
	}
	ctx.JSON(200, response)
}
