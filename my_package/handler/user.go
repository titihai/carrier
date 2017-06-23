package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
author : tiger
2017/5/31
用户模块相关操作
*/

// Login 用户登陆
func Login(c *gin.Context) {
	account := c.Param("account")
	password := c.Param("password")
	fmt.Print("tiger", account, password)
}
