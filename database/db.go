package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/nikhilesh-gupta/LibraryManagementSystem/models"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "books"
	username := "postgres" //change according to your own username
	password := "postgres" //change according to your own password
	args := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable password=" + password
	db, err := gorm.Open("postgres", args) //Provide which database you are using in the first argument (like here: postgres)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{}) // create table automatically in the db
	DB = db

}
