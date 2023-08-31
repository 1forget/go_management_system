package middleware

import (
	"GolandProjects/School-Management/utils"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strings"
)

func AuthMiddleware(whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path

		// 检查请求的路径是否在白名单中
		for _, allowedPath := range whitelist {
			if path == allowedPath {
				c.Next() // 跳过认证，继续执行下一个处理程序
				return
			}
		}
		//验证token
		tokenString := c.GetHeader("Authorization")
		//判断token是否已被动过期 即在黑名单中
		if exists := utils.IsIncludeExpiredToken(tokenString); exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		token, err := utils.GetJWTManager().VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		// 在这里 验证ip
		//执行Token验证逻辑
		if tokenString != "" && token.ClientIp == GetIP(c) {
			substr := path[18:23]
			//管理员验证token 管理员的Role字段为1
			if substr == "admin" && token.Role < 1 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			}
			c.Next() // 继续执行下一个处理程序
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		}
	}
}
func GetIP(c *gin.Context) string {
	// 尝试获取 X-Real-IP 头部
	realIP := c.GetHeader("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// 尝试获取 X-Forwarded-For 头部，获取第一个 IP 地址
	forwardedFor := c.GetHeader("X-Forwarded-For")
	if forwardedFor != "" {
		ips := strings.Split(forwardedFor, ",")
		for _, ip := range ips {
			// 清除可能的空格
			parsedIP := strings.TrimSpace(ip)
			if parsedIP != "" {
				return parsedIP
			}
		}
	}

	// 如果上述头部都不存在，使用请求的 RemoteAddr 字段
	// 但请注意，这可能包含代理服务器的地址而不是客户端的真实 IP
	remoteAddr := c.Request.RemoteAddr
	ip, _, _ := net.SplitHostPort(remoteAddr)
	return ip
}
