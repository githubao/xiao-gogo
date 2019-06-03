// gin
// author: baoqiang
// time: 2019-04-02 11:54
package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloGin() {
	e := gin.Default()

	e.Handle("GET", "/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"msg": "pong"})
	})

	_ = e.Run(":8080")
}
