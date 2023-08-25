package dao

import (
	"GolandProjects/School-Management/models"
	"GolandProjects/School-Management/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(admin models.Admin, clientIP string) string {
	// 验证密码
	db := utils.GetDB()

	loginPassword := admin.Password

	//db.First(&product, "Code = ?", "D43")
	//用户不存在
	db.First(&admin, "name=?", admin.Name)
	if admin.IsEmpty() {
		return "用户名或密码错误"
	}
	//使用 bcrypt 进行密码解密
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginPassword))
	if err != nil {
		return "用户名或密码错误"
	}
	fmt.Println("IP Address:" + clientIP)
	token, _ := utils.GetJWTManager().GenerateAdminToken(admin, clientIP)
	return token
}
func UpdateAdmin(admin models.Admin) error {

	db := utils.GetDB()
	if admin.Password != "" {
		//使用 bcrypt 进行密码哈希
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		admin.Password = string(hashedPassword)
	}

	result := db.Updates(&admin)
	//if err := db.Create(user).Error; err != nil {
	//	fmt.Println("插入失败", err)
	//}
	return result.Error
}

func GetAllUser() []models.User {
	db := utils.GetDB()
	users := make([]models.User, 0)
	//db.Find(&products,
	db.Find(&users)
	return users
}
