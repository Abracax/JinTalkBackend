package talk

import (
	"github.com/Abracax/JinTalkBackend/utils"
	"github.com/gin-gonic/gin"
)

func Random(c *gin.Context) {
	msg := ""
	exclamation := utils.GetRandFromConfEntry("exclamation")
	mid := utils.GetRandFromConfEntry("middle")
	politics := utils.GetRandFromConfEntry("politics")
	end := utils.GetRandFromConfEntry("other")

	msg += exclamation  + mid + politics + end
	c.JSON(200, gin.H{
		"status":  "success",
		"message": msg,
	})
}