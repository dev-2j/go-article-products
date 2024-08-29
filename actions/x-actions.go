package actions

import (
	"github.com/dev-2j/libaryx/validx"
	"github.com/gin-gonic/gin"
)

func XActions(c *gin.Context) {

	x := `a`
	y := []string{"a", "b", "c"}
	if validx.IsContains(y, x) {
		c.JSON(200, gin.H{
			"message": "x is empty",
		})
		return
	}

}
