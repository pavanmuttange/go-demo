package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error
	err = godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, "postgres", "postgres", "postgres", port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error in connecting database")
	}

	// DB, err = sql.Open("postgres", dsn)
	// if err != nil {
	// 	fmt.Println("error in connecting database", err.Error())
	// }
	fmt.Println("Database connected successfully")

}

func Init() {

	ConnectDB()
}
