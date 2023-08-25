// Package classification My cool service.
//
// the purpose of this service is to provide a
// mecahnism for experts to do their things.
//
//	Schemes: http
//	Host: localhost:8080
//	Version: 0.0.1
//	License: MIT http://opensource.org/licenses/MIT
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package controller

import (
	"GolandProjects/School-Management/dao"
	"GolandProjects/School-Management/middleware"
	"GolandProjects/School-Management/models"
	"GolandProjects/School-Management/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

type UserHandler struct{}

func (uh *UserHandler) SetupRouter(r *gin.Engine) {

	//白名单
	whitelist := []string{
		"/user/CreateUser",
		"/user/Login",
	}

	userGroup := r.Group("/user", middleware.AuthMiddleware(whitelist))
	// 在 userGroup 路由组内定义需要使用前缀的路由
	userGroup.POST("/Login", uh.Login)
	userGroup.PUT("/UpdateUser", uh.UpdateUser)
	userGroup.DELETE("/DeleteUser", uh.DeleteUser)
	userGroup.POST("/CreateUser", uh.CreateUser)
	userGroup.GET("/GetUser", uh.GetUser)
}

// curl -X POST -H "Content-Type: application/json" -d "{\"studentId\": 12, \"password\": \"pass12\"}" http://localhost:8080/user/CreateUser
// addUser
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	err := dao.AddUser(user)
	if err != nil {
		c.JSON(201, gin.H{"message": "Add Failure"})
		return
	}
	c.JSON(200, gin.H{"message": "User Add Success"})
}

// swagger:operation POST /user/Login user addUser
// ---
// summary: 登录
// description: 用于系统用户的登录
// parameters:
//   - name: studentId
//     in: body
//     description: 学号
//     type: int
//     required: true
//   - name: password
//     in: body
//     description: 密码
//     type: string
//     required: true

// curl -X POST -H "Content-Type: application/json" -d "{\"studentId\": 12, \"password\": \"pass12\"}" http://localhost:8080/user/Login
// login
func (uh *UserHandler) Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	userToken := dao.UserLogin(user, middleware.GetIP(c))

	fmt.Println(userToken)

	if userToken == "用户名或密码错误" {
		c.JSON(200, gin.H{"message": "Login Failure"})
		return
	}
	c.JSON(200, gin.H{"message": "Login Success", "JWT": userToken})
}

// curl -X PUT -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NiwiY2xpZW50SVAiOiIxMjcuMC4wLjEiLCJjcmVhdGVkQXQiOiIyMDIzLTA4LTI0VDEzOjQ5OjU4LjcyNjQwOCswODowMCIsImV4cCI6MTY5Mjk0NDUzNywiZ3JhZGUiOiIiLCJwYXNzd29yZCI6IiQyYSQxMCRaYWNIWFJ2MzNFQ1dyMWlGSjVheHdPREN3TVF0NmlBd3lJSnBhUTZJMXViVy5YNTFndmhmTyIsInN0dWRlbnRfaWQiOjEyLCJ1c2VybmFtZSI6IiJ9.YaJVABRHffodGi72dETOuw1hap1NCDP5IM8T-Ga5zzc" -d "{\"studentId\": 12, \"password\": \"pass12\",\"grade\": \"pass12\"}" http://localhost:8080/user/UpdateUser
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	//验证token
	tokenString := c.GetHeader("Authorization")

	token, err2 := utils.GetJWTManager().VerifyToken(tokenString)
	if err2 != nil {
		c.JSON(401, "Unauthorized")
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to read request body"})
		return
	}
	var user models.User
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	user.ID = token.ID
	fmt.Println(user)
	updateErr := dao.UpdateUser(user)
	if updateErr != nil {
		c.JSON(500, "Update Failure")
	}
	//重新签发token
	ip := middleware.GetIP(c)
	userToken := dao.UserLogin(user, ip)
	//使旧的token过期
	utils.SetExpiredToken(tokenString, ip)
	c.JSON(200, gin.H{"message": "Update Success", "JWT": userToken})

}

// curl -X DELETE -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NywiY2xpZW50SVAiOiIxMjcuMC4wLjEiLCJjcmVhdGVkQXQiOiIyMDIzLTA4LTI0VDE0OjMyOjU2LjU1NTY3MiswODowMCIsImV4cCI6MTY5Mjk0NTQ4MCwiZ3JhZGUiOiIiLCJwYXNzd29yZCI6IiQyYSQxMCR2V1BFMmJINS9yLjJSS3V5M0svNy8uSGlrUndnVFUzR29rTEl1UnRaV3k0YTZBblFFNW53bSIsInN0dWRlbnRfaWQiOjEzLCJ1c2VybmFtZSI6IiJ9.tZJN4rM4dnNtQSGgX-nWgSHlOPBMZilwxQRowexZmdY"  http://localhost:8080/user/DeleteUser
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	//验证token
	tokenString := c.GetHeader("Authorization")

	token, err2 := utils.GetJWTManager().VerifyToken(tokenString)
	if err2 != nil {
		c.JSON(401, "Unauthorized")
		return
	}

	err2 = dao.DeleteUserByUserId(token.ID)
	if err2 != nil {
		c.JSON(500, "Delete Failure")
		return
	}
	ip := middleware.GetIP(c)
	utils.SetExpiredToken(tokenString, ip)
	c.JSON(200, "Delete Success")

}

// curl -X GET -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NywiY2xpZW50SVAiOiIxMjcuMC4wLjEiLCJjcmVhdGVkQXQiOiIyMDIzLTA4LTI0VDE0OjMyOjU2LjU1NTY3MiswODowMCIsImV4cCI6MTY5Mjk0NTQ4MCwiZ3JhZGUiOiIiLCJwYXNzd29yZCI6IiQyYSQxMCR2V1BFMmJINS9yLjJSS3V5M0svNy8uSGlrUndnVFUzR29rTEl1UnRaV3k0YTZBblFFNW53bSIsInN0dWRlbnRfaWQiOjEzLCJ1c2VybmFtZSI6IiJ9.tZJN4rM4dnNtQSGgX-nWgSHlOPBMZilwxQRowexZmdY"  http://localhost:8080/user/GetUser
func (uh *UserHandler) GetUser(c *gin.Context) {
	//验证token
	tokenString := c.GetHeader("Authorization")

	token, err2 := utils.GetJWTManager().VerifyToken(tokenString)
	if err2 != nil {
		c.JSON(401, "Unauthorized")
		return
	}
	c.JSON(200, gin.H{"message": token})

}
