package db

import (
	"fmt"
	"go-redis-mysql_userAuth/pkg/config"
	"go-redis-mysql_userAuth/pkg/domain"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySql(config *config.Config) *gorm.DB {
	fmt.Println(config.SqlDSN)
	db, err := gorm.Open(mysql.Open(config.SqlDSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect MySQL")

	}
	log.Println("Successfully connected to MySQL")
	return db
}
func SyncDB(db *gorm.DB) {
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		panic("Failed to sync 'User'")
	}
}
