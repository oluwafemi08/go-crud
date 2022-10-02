package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oluwafemi08/go-app/pkg/models"
	"github.com/oluwafemi08/go-app/pkg/utils"
)

var NewUser models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	getUsers := models.GetAllUsers()
	res, _ := json.Marshal(getUsers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, _ := models.GetUserByID(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	user := CreateUser.CreateUser()
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateUser := &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetails, db := models.GetUserByID(ID)

	if updateUser.FirstName != "" {
		userDetails.FirstName = updateUser.FirstName
	}

	if updateUser.LastName != "" {
		userDetails.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.LastName
	}

	if updateUser.Message != "" {
		userDetails.Message = updateUser.Message
	}

	if updateUser.Address != " " {
		userDetails.Address = updateUser.Address
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
