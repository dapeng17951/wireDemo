/*
 * @Description:
 * @Author: shangke
 * @Date: 2024/10/26
 * @LastEditTime: 2024/10/26
 * @LastEditors: shangke
 */
package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowAccount godoc
// @Summary      测试2-扫描其它package下的方法
// @Description  输出json数据
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /ping2 [get]
func PingFun2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong2",
	})
}
