package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/huangzhuo492008824/go-gin-example/pkg/e"
	"github.com/huangzhuo492008824/go-gin-example/pkg/util"
)

func JWT(c *gin.Context) {
	var code int
	var data interface{}

	code = e.SUCCESS
	token := c.GetHeader("token")
	// if c.Request.Header.Get("token"):
	// token := c.Request.Header["token"][0]
	if token == "" {
		code = e.INVALID_PARAMS
	} else {
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}
	}
	if code != e.SUCCESS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		c.Abort()
	}
	c.Next()
	return
}
