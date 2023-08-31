package main

import (
	"GolandProjects/School-Management/controller"
	"GolandProjects/School-Management/middleware"
	"GolandProjects/School-Management/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	//连接数据库
	utils.SetupDB()
	//连接redis
	utils.SetUpRedis()
	//初始化JWT
	secretKey := "weomssaxiao148"
	utils.NewJWTManager(secretKey)

	//设置白名单
	whitelist := []string{
		"/schoolManagement/user/createUser",
		"/schoolManagement/user/login",
	}

	r := gin.Default()
	//cors
	middleware.ServeCors(r)
	r.Use(middleware.AuthMiddleware(whitelist))
	// 创建处理器实例
	adminHandler := &controller.AdminHandler{}
	userHandler := &controller.UserHandler{}

	// 调用 SetupRouter 方法设置路由和处理逻辑
	adminHandler.SetupRouter(r)
	userHandler.SetupRouter(r)

	r.Run(":8080")

	//最后关闭数据库连接
	db := utils.GetDB()
	s, _ := db.DB()
	defer s.Close()

	client := utils.GetRedisClient()
	defer client.Close()
}
