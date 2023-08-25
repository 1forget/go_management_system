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

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Name
	claims["ID"] = user.ID
	claims["password"] = user.Password
	claims["student_id"] = user.StudentId
	claims["grade"] = user.Grade
	claims["createdAt"] = user.CreatedAt
	claims["clientIP"] = clientIP
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jm.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (jm *JWTManager) GenerateAdminToken(admin models.Admin, clientIP string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = admin.Name
	claims["ID"] = admin.ID
	claims["password"] = admin.Password
	claims["createdAt"] = admin.CreatedAt
	claims["clientIP"] = clientIP
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

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		parsedTime, _ := time.Parse(time.RFC3339, claims["createdAt"].(string))
		//ID uint, name string,CreatedAt time.Time, studentID uint, grade string,clientIP string
		user2 := dto.NewUserDTO(
			uint(claims["ID"].(float64)),
			claims["username"].(string),
			parsedTime,
			uint(claims["student_id"].(float64)),
			claims["grade"].(string),
			claims["clientIP"].(string),
		)
		return *user2, nil
	} else {
		return user, fmt.Errorf("Invalid JWT")
	}
}

func (jm *JWTManager) VerifyAdminToken(tokenString string) (dto.AdminDTO, error) {

	var admin dto.AdminDTO
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jm.secretKey, nil
	})

	if err != nil {
		return admin, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		parsedTime, _ := time.Parse(time.RFC3339, claims["createdAt"].(string))
		//ID uint, name string,CreatedAt time.Time, studentID uint, grade string,clientIP string
		admin2 := dto.NewAdminDTO(

			uint(claims["ID"].(float64)),
			claims["name"].(string),
			parsedTime,

			claims["password"].(string),
			claims["clientIP"].(string),
		)
		return *admin2, nil
	} else {
		return admin, fmt.Errorf("Invalid JWT")
	}
}
