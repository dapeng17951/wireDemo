/*
 * @Description:
 * @Author: shangke
 * @Date: 2024/10/26
 * @LastEditTime: 2024/10/26
 * @LastEditors: shangke
 */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"net/http"
	"wireDemo/docs"
	"wireDemo/handle"
)

func main() {
	docs.SwaggerInfo.Version = "v2.0"
	r := gin.Default()
	r.GET("/ping", pingFun)
	r.GET("/ping2", handle.PingFun2)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

// ShowAccount godoc
// @Summary      测试
// @Description  输出json数据
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /ping [get]
func pingFun(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
