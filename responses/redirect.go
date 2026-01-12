package responses

import "github.com/gin-gonic/gin"

func Redirect(c *gin.Context, location string) {
	c.Redirect(307, location)
}
