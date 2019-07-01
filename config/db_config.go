package config

import(
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"fmt"
	"cart/models"
	"log"
	"github.com/joho/godotenv"
)

var err error

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(
		&models.User{},
	)

	return db
}