package main

import (
	"go-redis-mysql_userAuth/pkg/config"
	"go-redis-mysql_userAuth/pkg/db"
	"go-redis-mysql_userAuth/pkg/handler"
	"go-redis-mysql_userAuth/pkg/repository"
	"go-redis-mysql_userAuth/pkg/routes"
	"go-redis-mysql_userAuth/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	sdb := db.ConnectMySql(&c)
	db.SyncDB(sdb)
	rdb := db.ConnectRedis(&c)
	var (
		userRepo    repository.UserRepo = repository.NewUserRepo(sdb, rdb)
		userService usecase.UserUseCase = usecase.NewUserUseCase(userRepo)
		userHandler handler.UserHandler = *handler.NewUserHandler(userService)
	)
	r := gin.Default()
	r.Use(gin.Logger())
	routes.UserRoutes(r, userHandler)
	r.Run(":" + c.Port)
}
