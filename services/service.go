package services

import (
	"gomod/database"
	"gomod/models"
)

func CreateUser(user *models.User) {
	database.DB.Create(&user)
}

func DeleteUser(deletedUser *models.User) {
	database.DB.Delete(&deletedUser)
}

func UpdateUser(updatedUser *models.User) {
	database.DB.Save(&updatedUser)
}

func GetUser(user *models.User) {
	database.DB.First(&user)
}

func GetUsers(users *[]models.User) {
	database.DB.Find(&users)
}
