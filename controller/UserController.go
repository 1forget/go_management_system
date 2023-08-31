package controller

import (
	"GolandProjects/School-Management/bean"
	"GolandProjects/School-Management/dao"
	"GolandProjects/School-Management/middleware"
	"GolandProjects/School-Management/models"
	"GolandProjects/School-Management/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type UserHandler struct{}

func (uh *UserHandler) SetupRouter(r *gin.Engine) {

	userGroup := r.Group("/schoolManagement/user")
	// 在 userGroup 路由组内定义需要使用前缀的路由
	userGroup.POST("/login", uh.Login)
	userGroup.PUT("/updateUser", uh.UpdateUser)
	userGroup.DELETE("/deleteUser", uh.DeleteUser)
	userGroup.POST("/createUser", uh.CreateUser)
	userGroup.GET("/getUser", uh.GetUser)
}

// curl -X POST -H "Content-Type: application/json" -d "{\"studentId\": 12, \"password\": \"pass12\"}" http://localhost:8080/user/CreateUser
// addUser
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		bean.ResponseError(c, 400, "Invalid JSON")
		return
	}
	err := dao.AddUser(user)
	if err != nil {
		bean.ResponseError(c, 500, "Add Failure")
		return
	}
	bean.ResponseSuccess(c, "User Add Success")
}

// curl -X POST -H "Content-Type: application/json" -d "{\"studentId\": 12, \"password\": \"pass12\"}" http://localhost:8080/user/Login
// login
func (uh *UserHandler) Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		bean.ResponseError(c, 400, "Invalid JSON")
		return
	}

	userToken := dao.UserLogin(user, middleware.GetIP(c))

	fmt.Println(userToken)

	if userToken == "用户名或密码错误" {
		bean.ResponseError(c, 403, "用户名或密码错误")
		return
	}
	bean.ResponseWithToken(c, "login success", userToken)
}

// curl -X PUT -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NiwiY2xpZW50SVAiOiIxMjcuMC4wLjEiLCJjcmVhdGVkQXQiOiIyMDIzLTA4LTI0VDEzOjQ5OjU4LjcyNjQwOCswODowMCIsImV4cCI6MTY5Mjk0NDUzNywiZ3JhZGUiOiIiLCJwYXNzd29yZCI6IiQyYSQxMCRaYWNIWFJ2MzNFQ1dyMWlGSjVheHdPREN3TVF0NmlBd3lJSnBhUTZJMXViVy5YNTFndmhmTyIsInN0dWRlbnRfaWQiOjEyLCJ1c2VybmFtZSI6IiJ9.YaJVABRHffodGi72dETOuw1hap1NCDP5IM8T-Ga5zzc" -d "{\"studentId\": 12, \"password\": \"pass12\",\"grade\": \"pass12\"}" http://localhost:8080/user/UpdateUser
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	//验证token 并取出token里的值
	tokenString := c.GetHeader("Authorization")

	token, err2 := utils.GetJWTManager().VerifyToken(tokenString)
	if err2 != nil {
		bean.ResponseError(c, 401, "Unauthorized")
		return
	}
	//将请求体中的json数据映射到userUpdateInfo中
	var userUpdateInfo bean.UpdateUserInfoRequest
	if err := c.ShouldBindJSON(&userUpdateInfo); err != nil {
		bean.ResponseError(c, 400, "Invalid JSON")
		return
	}
	//若改变密码
	if userUpdateInfo.Password != "" {
		//验证旧密码 是否有效
		userExists := dao.VerifyUser(userUpdateInfo.OldPassword, token.StudentId)
		if !userExists {
			bean.ResponseError(c, 400, "Wrong Old PassWord")
			return
		}
	}
	//将请求体中的user映射到实体类中的user中
	var user models.User
	err := mapstructure.Decode(userUpdateInfo, &user)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.ID = token.ID
	user.StudentId = token.StudentId

	updateErr := dao.UpdateUser(user)
	fmt.Println(user.ToString())
	if updateErr != nil {
		bean.ResponseError(c, 500, "Update Failure")
	}

	//token中保存了name id和studentId（不可改）
	//只有在改变了name才需要重新签发token
	userToken := ""
	if user.Name != "" {
		ip := middleware.GetIP(c)
		userToken, _ = utils.GetJWTManager().GenerateToken(user, ip)
		//使旧的token过期
		utils.SetExpiredToken(tokenString, ip)
	}
	//判断是否要重新生成token
	if userToken != "" {
		bean.ResponseWithToken(c, "update token", userToken)
	} else {
		bean.ResponseSuccess(c, "update success")
	}

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
	var deleteRequest bean.UserDeleteRequest
	if err := c.ShouldBindJSON(&deleteRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//验证用户密码是否正确
	userExists := dao.VerifyUser(deleteRequest.Password, token.StudentId)
	if !userExists {
		bean.ResponseError(c, 401, "Wrong Old PassWord")
		return
	}

	err2 = dao.DeleteUserByUserId(token.ID)
	if err2 != nil {
		bean.ResponseError(c, 500, "Delete Failure")

		return
	}
	ip := middleware.GetIP(c)
	utils.SetExpiredToken(tokenString, ip)

	bean.ResponseSuccess(c, "Delete Success")
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
	user := dao.GetUserById(token.ID)
	user.Password = ""
	userJsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	bean.ResponseSuccess(c, string(userJsonData))

}
