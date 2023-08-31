package dao

import (
	"GolandProjects/School-Management/models"
	"GolandProjects/School-Management/utils"
)

func GetAllUser() []models.User {
	db := utils.GetDB()
	users := make([]models.User, 0)
	//db.Find(&products,
	db.Find(&users)
	return users
}
