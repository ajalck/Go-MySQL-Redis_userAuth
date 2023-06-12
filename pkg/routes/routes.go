package routes

import (
	"go-redis-mysql_userAuth/pkg/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, handler handler.UserHandler) {
	router.POST("/register", handler.CreateUser)
	router.POST("/login", handler.UserLogin)
}
