package controllers

import (
	userModules "goApiPractice/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register User
func Register(c *gin.Context) {
	user := new(userModules.User)            // 宣告一個User變數空間
	if err := c.BindJSON(user); err != nil { // 將資料以json格式解析並存入user;err接收錯誤訊息
		c.JSON(http.StatusBadRequest, gin.H{ // 直接回覆錯誤
			"code":    0,
			"message": "Register fail",
		})
		return // 程式斷點 同break; (但go沒有break
	}

	result, err := user.Register()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"message": "Register fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": result,
	})
}
