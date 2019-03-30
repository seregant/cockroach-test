package main

import (
	"strings"

	gin "github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/seregant/cockroach-test/config"
	"github.com/seregant/cockroach-test/controllers"
)

var conf = config.SetConfig()

var key = conf.SrvKey

func middleware() gin.HandlerFunc {
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

func main() {
	jabatanController := new(controllers.Jabatan)

	router := gin.Default()
	router.Use(middleware())

	jabatan := router.Group("/jabatan")
	{
		jabatan.GET("/", jabatanController.GetAllJabatan)
		jabatan.POST("/update/:id_jabatan", jabatanController.UpdateJabatan)
		jabatan.GET("/update/:id_jabatan", jabatanController.UpdateJabatan)
	}

	router.Run(":" + conf.HttpPort)
}
