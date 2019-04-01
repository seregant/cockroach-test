package middleware

import (
	"strings"

	gin "github.com/gin-gonic/gin"
	"github.com/seregant/cockroach-test/config"
)

var conf = config.SetConfig()
var key = conf.SrvKey

func ServiceAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer != "" {
			strSplit := strings.Split(bearer, " ")
			if strSplit[0] == "Bearer" && strSplit[1] != "" {
				if strSplit[1] == key {
					c.Next()
				} else {
					c.AbortWithStatus(401)
					c.JSON(401, gin.H{
						"status":  401,
						"message": "unauthorized",
					})
				}
			} else {
				c.AbortWithStatus(401)
				c.JSON(401, gin.H{
					"status":  401,
					"message": "unauthorized",
				})
			}
		} else {
			c.AbortWithStatus(401)
			c.JSON(401, gin.H{
				"status":  401,
				"message": "unauthorized",
			})
		}
	}
}
