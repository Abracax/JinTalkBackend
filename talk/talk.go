package talk

import (
	"github.com/gin-gonic/gin"
)

func Talk(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		Random(c)
		return
	}

	msg, ok := ParseName(name)
	if ok {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": msg,
			"name":    name,
		})
		return
	}

	msg += "，"

	greeting := GetRandFromConfEntry("greetings")
	mid := GetRandFromConfEntry("middle")
	politics := GetRandFromConfEntry("politics")

	msg += greeting + mid + politics

	c.JSON(200, gin.H{
		"status":  "success",
		"message": msg,
		"name":    name,
	})
}

func Random(c *gin.Context) {
	msg := ""
	exclamation := GetRandFromConfEntry("exclamation")
	mid := GetRandFromConfEntry("middle")
	politics := GetRandFromConfEntry("politics")
	end := GetRandFromConfEntry("other")

	msg += exclamation  + mid + politics + end
	c.JSON(200, gin.H{
		"status":  "success",
		"message": msg,
	})
}
