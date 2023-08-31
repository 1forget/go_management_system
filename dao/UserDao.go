package dao

import (
	"GolandProjects/School-Management/models"
	"GolandProjects/School-Management/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(user models.User) error {

	db := utils.GetDB()
	//使用 bcrypt 进行密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)

	result := db.Create(&user)
	//if err := db.Create(user).Error; err != nil {
	//	fmt.Println("插入失败", err)
	//}
	return result.Error
}
func UpdateUser(user models.User) error {

	db := utils.GetDB()
	if user.Password != "" {
		//使用 bcrypt 进行密码哈希
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		user.Password = string(hashedPassword)
	}

	result := db.Updates(&user)
	//if err := db.Create(user).Error; err != nil {
	//	fmt.Println("插入失败", err)
	//}
	return result.Error
}

func UserLogin(user models.User, clientIP string) string {
	// 验证密码
	db := utils.GetDB()

	loginPassword := user.Password

	//db.First(&product, "Code = ?", "D43")
	//用户不存在
	db.First(&user, "student_id=?", user.StudentId)
	if user.IsEmpty() {
		return "用户名或密码错误"
	}
	//使用 bcrypt 进行密码解密
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPassword))
	if err != nil {
		return "用户名或密码错误"
	}
	fmt.Println("IP Address:" + clientIP)
	token, _ := utils.GetJWTManager().GenerateToken(user, clientIP)

	return token
}
func VerifyUser(password string, studentId uint) bool {
	db := utils.GetDB()
	//用户不存在
	var user models.User
	db.First(&user, "student_id=?", studentId)
	if user.IsEmpty() {
		return false
	}
	//使用 bcrypt 进行密码解密  第二个字符是待验证的密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
func DeleteUserByUserId(userID uint) error {
	db := utils.GetDB()
	var user models.User
	result := db.Delete(&user, userID)
	return result.Error
}
func GetUserById(id uint) models.User {
	db := utils.GetDB()
	var user models.User
	db.First(&user, "id=?", id)
	return user
}
