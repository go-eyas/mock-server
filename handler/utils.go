package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	origin := c.Request.Header["Origin"]
	if len(origin) > 0 {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin[0])
	}
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Max-Age", "3600")

	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	} else {
		c.Next()
	}
}

func APIHandler(c *gin.Context) {
	c.Next()
	status := c.Writer.Status()
	if status == 404 {
		method := c.Request.Method
		api := GetAPI(method, c.Request.RequestURI)
		if api != nil {
			c.JSON(200, api)
		}
	}

}

func getMethodAndPath(uri string) (string, string) {
	parsePath := strings.Split(uri, " ")
	var method, path string
	switch len(parsePath) {
	case 1:
		method = "any"
		path = parsePath[0]
	case 2:
		method = parsePath[0]
		path = parsePath[1]
	}
	return method, path
}
