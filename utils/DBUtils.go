package utils

import (
	"GolandProjects/School-Management/models"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Configuration struct {
	Database struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		DBName       string `yaml:"dbname"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"database"`
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
	}
}

var DB *gorm.DB // 包级别的变量，用于存储连接池

func SetUpTable() {
	var config Configuration
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&models.User{})
}

func SetupDB() {
	data, err := os.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.Port,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can not open database")
	}
	// 配置连接池参数
	sqlDB, err := DB.DB()
	if err != nil {
		panic("can not open database pool")
	}
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns) // 设置最大打开连接数
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns) // 设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Hour)                 // 设置连接的最大生存时间
}

func GetDB() *gorm.DB {
	if DB == nil {
		panic("can not connect DB")
	}
	return DB
}
