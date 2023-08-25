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

type AdminHandler struct{}

func (uh *AdminHandler) SetupRouter(r *gin.Engine) {
	adminGroup := r.Group("/admin")
	adminGroup.POST("/login", uh.login)
	adminGroup.PUT("/update", uh.update)
	adminGroup.GET("/getAllUser", uh.getAllUser)
}

// curl -X POST -H "Content-Type: application/json" -d "{\"name\":\"admin111\" , \"password\": \"pass12\"}" http://localhost:8080/admin/login
func (uh *AdminHandler) login(c *gin.Context) {
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	userToken := dao.AdminLogin(admin, middleware.GetIP(c))

	fmt.Println(userToken)

	if userToken == "用户名或密码错误" {
		c.JSON(200, gin.H{"message": "Login Failure"})
		return
	}
	c.JSON(200, gin.H{"message": "Login Success", "JWT": userToken})
}

// curl -X PUT -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiY2xpZW50SVAiOiIxMjcuMC4wLjEiLCJjcmVhdGVkQXQiOiIyMDIzLTA4LTI0VDE1OjE5OjQ0LjE4NzIyNSswODowMCIsImV4cCI6MTY5Mjk2OTE5OSwibmFtZSI6ImFkbWluMTExIiwicGFzc3dvcmQiOiIkMmEkMTAkNmNOVnJrbURSQWYwZmNUMTV3SVlsLkVzQjFGTHFUVjFIS3VxclJYODA0UFE5bFpYVGdGenkifQ.wOoFe2CcFcWhil0xU4GT1I7rVa_cVaEwLH3EYHw2vuo" -d "{\"ID\": 7, \"password\": \"pass12\",\"grade\": \"admintest\"}" http://localhost:8080/admin/update
func (uh *AdminHandler) update(c *gin.Context) {
	//验证token
	tokenString := c.GetHeader("Authorization")

	token, err2 := utils.GetJWTManager().VerifyAdminToken(tokenString)
	if tokenString == "" || err2 != nil {
		c.JSON(401, "Unauthorized")
		return
	}
	//对比IP
	if token.ClientIp != middleware.GetIP(c) {
		c.JSON(402, "wrong IP")
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

	fmt.Println(user)
	updateErr := dao.UpdateUser(user)
	if updateErr != nil {
		c.JSON(500, "Update Failure")
	}

	c.JSON(200, "Update Success")
}

// curl -X GET -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiY2xpZW50SVAiOiIxMjcuMC4wLjEiLCJjcmVhdGVkQXQiOiIyMDIzLTA4LTI0VDE1OjE5OjQ0LjE4NzIyNSswODowMCIsImV4cCI6MTY5Mjk2OTE5OSwibmFtZSI6ImFkbWluMTExIiwicGFzc3dvcmQiOiIkMmEkMTAkNmNOVnJrbURSQWYwZmNUMTV3SVlsLkVzQjFGTHFUVjFIS3VxclJYODA0UFE5bFpYVGdGenkifQ.wOoFe2CcFcWhil0xU4GT1I7rVa_cVaEwLH3EYHw2vuo" http://localhost:8080/admin/getAllUser
func (uh *AdminHandler) getAllUser(c *gin.Context) {

	//验证token
	tokenString := c.GetHeader("Authorization")

	token, err2 := utils.GetJWTManager().VerifyAdminToken(tokenString)
	if tokenString == "" || err2 != nil {
		c.JSON(401, "Unauthorized")
		return
	}
	//对比IP
	if token.ClientIp != middleware.GetIP(c) {
		c.JSON(402, "wrong IP")
		return
	}
	users := dao.GetAllUser()
	//遍历user将password为空
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(200, users)
}
