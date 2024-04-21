package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	var (
		// db config
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbUsername = os.Getenv("DB_USERNAME")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
	)

	// connect db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		fmt.Println("dsn: ", dsn)
		panic(err)
	}

	fmt.Println("Connected to database")

	//db.AutoMigrate(&task.Task{})

	//fmt.Println("Database Migrated")

	return db, err

}
