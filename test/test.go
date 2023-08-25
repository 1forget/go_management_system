package main

import (
	"GolandProjects/School-Management/dao"
	"GolandProjects/School-Management/models"
	"GolandProjects/School-Management/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func testAddUser() {

	// 创建一个 User 结构体实例
	user := models.User{
		ID:        1,
		CreatedAt: time.Now(),
		Name:      "user123",
		StudentId: 12345,
		Grade:     "A",
		Password:  "pass123",
	}
	err2 := dao.AddUser(user)

	if err2 != nil {
		panic(err2)
	}
}
func testPasswordEncrypt() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("pa123"), bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))
	hashed, _ := bcrypt.GenerateFromPassword([]byte("pa123"), bcrypt.DefaultCost)
	fmt.Println(string(hashed))

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("pa123"))
	fmt.Println(err)
	err2 := bcrypt.CompareHashAndPassword([]byte(hashed), []byte("pa123"))
	fmt.Println(err2)
}

func testJWT() {
	user := models.User{
		ID:        1,
		CreatedAt: time.Now(),
		Name:      "user123",
		StudentId: 12345,
		Grade:     "A",
		Password:  "pass123",
	}

	secretKey := "weomssaxiao148"
	jwtManager := utils.NewJWTManager(secretKey)
	token, err := jwtManager.GenerateToken(user, "127.0.0.1")
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated JWT:", token)

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NywiY2xpZW50SVAiOiIxMjcuMC4wLjEiLCJjcmVhdGVkQXQiOiIyMDIzLTA4LTI0VDE0OjMyOjU2LjU1NTY3MiswODowMCIsImV4cCI6MTY5Mjk0NTkxOCwiZ3JhZGUiOiIiLCJwYXNzd29yZCI6IiQyYSQxMCR2V1BFMmJINS9yLjJSS3V5M0svNy8uSGlrUndnVFUzR29rTEl1UnRaV3k0YTZBblFFNW53bSIsInN0dWRlbnRfaWQiOjEzLCJ1c2VybmFtZSI6IiJ9.E00bV514RNOWKbYASxbqQoUQK6iLf_ujZ0ud7Cq3DKs"

	tokenUser, err := jwtManager.VerifyToken(token)

	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}

	fmt.Println("Valid JWT for user:", tokenUser)
}
func getAllUser() {
	db := utils.GetDB()
	users := make([]models.User, 0)
	//db.Find(&products,
	db.Find(&users)
	fmt.Println(users)
}
func main() {
	utils.SetupDB()
	//testAddUser()
	//testPasswordEncrypt()
	//testJWT()
	getAllUser()
}
