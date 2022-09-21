package connect

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error
var count = 0

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	Db, err = connect(dsn)

	if err != nil {
		log.Fatalln(dsn + "Database can't connect")
	}
}

func connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Not ready. Retry connecting...")
		time.Sleep(time.Second)
		count++

		log.Println(count)
		if count > 30 {
			panic(err)
		}

		return connect(dsn)
	}

	log.Println("Database connecting is successfully")
	return db, nil
}
