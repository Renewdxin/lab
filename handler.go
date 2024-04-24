package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandlerPort interface {
	Login(c *gin.Context)
}

type UserHandlerAdapter struct {
	dao UserDaoPort
}

func NewUserHandlerAdapter(dao UserDaoPort) UserHandlerAdapter {
	return UserHandlerAdapter{dao}
}

func (handler UserHandlerAdapter) Login(c *gin.Context) {
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
