package middleware

import (
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	// Do something before the endpoint is called
	c.Next()
	// Do something after the endpoint is called
}
