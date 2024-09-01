package main

import "os"
import "log"
import "gorm.io/driver/mysql"
import "gorm.io/gorm"
import "fmt"

func Connect() (*gorm.DB, error) {

	rooter := fmt.Sprintf(
		"root:%s@tcp(%s:3306)/%s", 
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("CONTAINER_MYSQL"),
		os.Getenv("MYSQL_DATABASE"),
	)
	
	database, err := gorm.Open(mysql.Open(rooter), &gorm.Config{})
	if err != nil {
		log.Println(err);
	}
	log.Println("Connect !")
	return database, err
}

func main(){
	database, err := Connect()

	if err != nil {
		log.Println(err)
	}

	// Create Table
	type Mahasiswas struct {
		gorm.Model
		UserMahasiswa string
	}

	errmit := database.AutoMigrate(&Mahasiswas{})

	if errmit != nil {
		log.Println("Migrate !")
	}
}