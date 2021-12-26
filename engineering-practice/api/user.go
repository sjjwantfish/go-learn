package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sjjwantfish/go-learn/engineering-practice/internal/biz/user"
	"github.com/sjjwantfish/go-learn/engineering-practice/internal/service"
	"time"
)


func CreateUser(c *gin.Context)  {
	userInfo := service.UserInfoDto{}
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(10001, gin.H{"error": "参数错误"})
		return
	}
	do := user.UserInfoDo{}
	do.Tel = userInfo.Tel
	do.Passwd = userInfo.Passwd
	do.Name = userInfo.Name
	do.CreateTime = time.Now()
	do.Create()
	c.JSON(200, gin.H{"msg": "创建成功"})
	return
}
