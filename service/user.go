package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserName(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户id不能为空",
		})
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	switch userId {
	case 44:
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Lewis Hamilton",
		})
	case 3:
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Daniel Ricardo",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "查无此人",
		})
		return
	}
}
