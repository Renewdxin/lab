package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var userHandler UserHandlerAdapter

func LoadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}
}

func initSQL() {
	sqlDB, err := sql.Open(os.Getenv("DRIVER_NAME"), os.Getenv("DRIVER_SOURCE_NAME"))
	if err != nil {
		return
	}

	DB, err = gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		return
	}
}

func InitializeRouter() {
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:63342"} // 允许的前端域
	r.Use(cors.New(config))

	// 1. 题目一，正常登录
	// 2. 在用户正常输入账户密码，无所谓是否正确时，便记录了用户的ip地址
	r.POST("/normal/login", userHandler.NormalLogin)
	// 题目五：防火墙系统
	r.POST("/filter/login", userHandler.IPUserFilterMiddleware, userHandler.NormalLogin)
	// 题目二、爆破
	r.POST("/crack/login", userHandler.CrackLoginHandler)
	// 题目三、此链接显示单个用户的登录历史
	r.GET("/login-history/:userID", userHandler.getLoginHistory)
	// 题目四、显示用户是否触及关键词
	r.POST("/check-pattern", userHandler.patternHandler)
	// 题目四、用户发送后端接收信息
	r.POST("/send-message", userHandler.KeywordDetectionMiddleware([]string{"password", "secret"}), userHandler.sendMessageHandler)
	// 题目七：图形验证模拟
	r.POST("image/login", userHandler.ImageUnlockHandler)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
