package utils

import (
	"GolandProjects/School-Management/dto"
	"GolandProjects/School-Management/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTManager struct {
	secretKey []byte
}

var manager JWTManager

func NewJWTManager(secretKey string) *JWTManager {
	manager = JWTManager{
		secretKey: []byte(secretKey),
	}
	return &manager
	//return &JWTManager{
	//	secretKey: []byte(secretKey),
	//}
}
func GetJWTManager() *JWTManager {
	return &manager
}
func (jm *JWTManager) GenerateToken(user models.User, clientIP string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	//生成token token中只存放名字 学号 IP以及Id
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Name
	claims["ID"] = user.ID
	claims["student_id"] = user.StudentId
	claims["clientIP"] = clientIP
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jm.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jm *JWTManager) VerifyToken(tokenString string) (dto.UserDTO, error) {

	var user dto.UserDTO
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jm.secretKey, nil
	})

	if err != nil {
		return user, err
	}

	//从token中取出信息
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//parsedTime, _ := time.Parse(time.RFC3339, claims["createdAt"].(string))
		//ID uint, name string,CreatedAt time.Time, studentID uint, grade string,clientIP string
		user2 := dto.NewUserDTO(
			uint(claims["ID"].(float64)),
			claims["username"].(string),
			uint(claims["role"].(float64)),
			uint(claims["student_id"].(float64)),
			claims["clientIP"].(string),
		)
		return *user2, nil
	} else {
		return user, fmt.Errorf("Invalid JWT")
	}
}
