package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandlerPort interface {
	NormalLogin(c *gin.Context)
	CrackLoginHandler(c *gin.Context)
	getLoginHistory(c *gin.Context)
}

type UserHandlerAdapter struct {
	dao UserDaoPort
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
	DB.Table("event").Create(&loginEvent)

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
	result := DB.Where("username = ?", loginInfo.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 模拟字典破解
	if crackPassword(loginInfo.Username, loginInfo.Password) {
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
