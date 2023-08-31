package controller

import (
	"GolandProjects/School-Management/bean"
	"GolandProjects/School-Management/dao"
	"GolandProjects/School-Management/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type AdminHandler struct{}

func (uh *AdminHandler) SetupRouter(r *gin.Engine) {
	adminGroup := r.Group("/schoolManagement/admin")
	adminGroup.PUT("/updateUser", uh.update)
	adminGroup.GET("/getAllUser", uh.getAllUser)
}

// curl -X PUT -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiY2xpZW50SVAiOiIxMjcuMC4wLjEiLCJjcmVhdGVkQXQiOiIyMDIzLTA4LTI0VDE1OjE5OjQ0LjE4NzIyNSswODowMCIsImV4cCI6MTY5Mjk2OTE5OSwibmFtZSI6ImFkbWluMTExIiwicGFzc3dvcmQiOiIkMmEkMTAkNmNOVnJrbURSQWYwZmNUMTV3SVlsLkVzQjFGTHFUVjFIS3VxclJYODA0UFE5bFpYVGdGenkifQ.wOoFe2CcFcWhil0xU4GT1I7rVa_cVaEwLH3EYHw2vuo" -d "{\"ID\": 7, \"password\": \"pass12\",\"grade\": \"admintest\"}" http://localhost:8080/admin/update
func (uh *AdminHandler) update(c *gin.Context) {
	//读取请求体中的JSON数据
	var userUpdateInfo bean.AdminUpdateUserRequest

	if err := c.ShouldBindJSON(&userUpdateInfo); err != nil {
		fmt.Println(err)
		bean.ResponseError(c, 400, "Invalid JSON")
		return
	}
	var user models.User
	err := mapstructure.Decode(userUpdateInfo, &user)

	if err != nil {
		bean.ResponseError(c, 400, "Invalid JSON")
		return
	}

	//修改对应user
	updateErr := dao.UpdateUser(user)
	if updateErr != nil {
		bean.ResponseError(c, 500, "Update Failure")
		return
	}
	bean.ResponseSuccess(c, "Update Success")
}

// curl -X GET -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiY2xpZW50SVAiOiIxMjcuMC4wLjEiLCJjcmVhdGVkQXQiOiIyMDIzLTA4LTI0VDE1OjE5OjQ0LjE4NzIyNSswODowMCIsImV4cCI6MTY5Mjk2OTE5OSwibmFtZSI6ImFkbWluMTExIiwicGFzc3dvcmQiOiIkMmEkMTAkNmNOVnJrbURSQWYwZmNUMTV3SVlsLkVzQjFGTHFUVjFIS3VxclJYODA0UFE5bFpYVGdGenkifQ.wOoFe2CcFcWhil0xU4GT1I7rVa_cVaEwLH3EYHw2vuo" http://localhost:8080/admin/getAllUser
func (uh *AdminHandler) getAllUser(c *gin.Context) {

	//取出数据
	users := dao.GetAllUser()
	//遍历user将password为空
	for i := range users {
		users[i].Password = ""
	}
	jsonUserData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	bean.ResponseSuccess(c, string(jsonUserData))
}
