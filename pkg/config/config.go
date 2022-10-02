package config

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)


var (
	db *gorm.DB
)

func Connect() {

dialect := os.Getenv("DIALECT")
host := os.Getenv("HOST")
dbPORT := os.Getenv("PORT")
dbNAME := os.Getenv("NAME")
user := os.Getenv("USER")
password := os.Getenv("PASSWORD")

dbURI := fmt.Sprintf("host=%s user=%s dbNAME=%s dbPORT=%s password=%s sslmode=disable", host, user, dbNAME, dbPORT, password)
	

d, err := gorm.Open(dialect, dbURI)
	if err != nil {
	
	panic(err)
		
	}
	fmt.Println("Successfully connected to database!")
	db = d
	
}



func GetDB() *gorm.DB {
	return db
}
