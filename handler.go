package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandlerPort interface {
	NormalLogin(c *gin.Context)
	CrackLoginHandler(c *gin.Context)
	patternHandler(c *gin.Context)
	getLoginHistory(c *gin.Context)
	KeywordDetectionMiddleware(keywords []string) gin.HandlerFunc
	sendMessageHandler(c *gin.Context)
	ImageUnlockHandler(c *gin.Context)
}

type UserHandlerAdapter struct {
	dao UserDaoPort
}

var (
	blacklistedIPs   = []string{"192.168.1.4", "10.0.0.2"}
	blacklistedUsers = []string{"user1", "user2"}
)

type LoginRequest struct {
	Username string   `json:"username"`
	Pattern  []string `json:"pattern"`
}

var userPatterns = map[string][]string{
	"alice":   {"1,1", "1,2", "2,2", "2,3"},
	"bob":     {"2,1", "3,1", "3,2", "4,2"},
	"charlie": {"1,4", "2,4", "3,4", "4,4"},
}

func NewUserHandlerAdapter(dao UserDaoPort) UserHandlerAdapter {
	return UserHandlerAdapter{dao}
}

func (handler UserHandlerAdapter) NormalLogin(c *gin.Context) {
	var req struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	loginEvent := LoginEvent{
		UserID:    req.ID,
		LoginTime: time.Now(),
		IP:        c.ClientIP(),
	}
	if err := DB.Table("event").Create(&loginEvent).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}

	c.JSON(200, gin.H{"message": "Logged in successfully"})

	if err := handler.dao.Login(req.Password, req.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (handler UserHandlerAdapter) CrackLoginHandler(c *gin.Context) {
	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{}
	if err := DB.Where("username = ?", loginInfo.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 模拟字典破解
	if crackPassword(loginInfo.Password) {
		c.JSON(http.StatusOK, gin.H{"message": "Password cracked successfully"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password cracking failed"})
	}
}

func (handler UserHandlerAdapter) getLoginHistory(c *gin.Context) {
	userID := c.Param("userID") // 从URL路径获取用户ID
	var events []LoginEvent
	//TODO
	result := DB.Where("user_id = ?", userID).Find(&events)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No login history found"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (handler UserHandlerAdapter) patternHandler(c *gin.Context) {
	var input struct {
		Pattern []string `json:"pattern"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if checkPattern(input.Pattern) {
		c.JSON(200, gin.H{"message": "Pattern correct"})
	} else {
		c.JSON(403, gin.H{"message": "Pattern incorrect"})
	}
}

func (handler UserHandlerAdapter) IPUserFilterMiddleware(c *gin.Context) {
	user := c.PostForm("username")
	ip := c.ClientIP()

	if contains(blacklistedIPs, ip) || contains(blacklistedUsers, user) {
		c.AbortWithStatusJSON(403, gin.H{"error": "Access denied"})
		return
	}
	c.Next()
}

func (handler UserHandlerAdapter) KeywordDetectionMiddleware(keywords []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var message struct {
			Text string `json:"text"`
		}
		if err := c.BindJSON(&message); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		for _, keyword := range keywords {
			if strings.Contains(message.Text, keyword) {
				fmt.Println("Detected keyword:", message.Text) // 可以替换为其他记录行为
				break
			}
		}

		c.Next()
	}
}

// 使用中间件

func contains(list []string, item string) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
}

func (handler UserHandlerAdapter) sendMessageHandler(c *gin.Context) {
	var message struct {
		Text string `json:"text"`
	}

	// 绑定JSON体到message变量
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message format"})
		return
	}

	// 记录接收到的消息到日志，实际应用中可以替换为其他类型的处理，如存储到数据库
	log.Printf("Received message: %s", message.Text)

	// 响应客户端，确认消息已被接收
	c.JSON(http.StatusOK, gin.H{"status": "Received"})
}

func (handler UserHandlerAdapter) ImageUnlockHandler(c *gin.Context) {
	var request LoginRequest
	if err := c.BindJSON(&request); err == nil {
		expectedPattern, ok := userPatterns[request.Username]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "unknown user"})
			return
		}
		if len(request.Pattern) != len(expectedPattern) {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "invalid pattern length"})
			return
		}
		for i, point := range expectedPattern {
			if request.Pattern[i] != point {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "pattern does not match"})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "pattern verified"})
	}
}
