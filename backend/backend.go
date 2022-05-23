package backend

import "github.com/gin-gonic/gin"

func Run() {
	r := gin.Default()
	r.GET("/unbond", func(c *gin.Context) {
		var validator= c.Query("validator")
		c.JSON(200, gin.H{
			"message": getUnbondFromValidator(validator),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}