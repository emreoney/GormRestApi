package handlers

import (
	"encoding/json"
	"fmt"
	"gomod/models"
	"gomod/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	services.CreateUser(&user)

	data, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintf(w, string(data))
}

func HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	userId, err := strconv.Atoi(variables["id"])
	if err != nil {
		fmt.Println(err.Error())
	}
	var deletedUser models.User
	deletedUser.ID = uint(userId)

	services.DeleteUser(&deletedUser)

	responseMessage := models.ResponseInformation{"ID:" + variables["id"] + " has deleted"}
	responseMessageData, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprint(w, string(responseMessageData))

}

func HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	userID, err := strconv.Atoi(variables["id"])
	if err != nil {
		fmt.Println(err.Error())
	}

	var updatedUser models.User
	updatedUser.ID = uint(userID)

	json.NewDecoder(r.Body).Decode(&updatedUser)
	services.UpdateUser(&updatedUser)

	data, err := json.Marshal(updatedUser)
	fmt.Fprintf(w, string(data))
}

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	services.GetUsers(&users)

	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprint(w, string(data))
}

func HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	variables := mux.Vars(r)
	userId, err := strconv.Atoi(variables["id"])
	if err != nil {
		fmt.Println(err.Error())
	}
	user.ID = uint(userId)
	services.GetUser(&user)

	data, err := json.Marshal(&user)
	fmt.Fprintf(w, string(data))

}
