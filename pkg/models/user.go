package models

import (
	"github.com/jinzhu/gorm"
	"github.com/oluwafemi08/go-app/pkg/config"
	_"github.com/lib/pq"
)

var db *gorm.DB

type User struct {
	gorm.Model
	FirstName string `gorm:"" json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Message   string `json:"message"`
	Address   string `json:"address"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserByID(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("ID=?", Id).Delete(user)
	return user

}
