package middleware

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

func ServeCors(router *gin.Engine) {
	cfg := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Type"},
		AllowCredentials: true,
	}
	router.Use(cors.New(cfg))
}
