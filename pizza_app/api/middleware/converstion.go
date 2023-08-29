package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseFormFloat64(c *gin.Context, key string) float64 {
	strValue := c.PostForm(key)
	value, _ := strconv.ParseFloat(strValue, 64)
	return value
}

func ParseFormInt(c *gin.Context, key string) int {
	strValue := c.PostForm(key)
	value, _ := strconv.ParseInt(strValue, 10, 0) // Assuming 32-bit int is enough
	return int(value)
}
