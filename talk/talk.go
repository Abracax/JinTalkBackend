package talk

import (
	"github.com/Abracax/JinTalkBackend/utils"
	"github.com/gin-gonic/gin"
)

func Talk(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		Random(c)
		return
	}

	msg, ok := parseName(name)
	if ok {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": msg,
			"name":    name,
		})
		return
	}

	msg += "，"

	greeting := utils.GetRandFromConfEntry("greetings")
	mid := utils.GetRandFromConfEntry("middle")
	politics := utils.GetRandFromConfEntry("politics")

	msg += greeting + mid + politics

	c.JSON(200, gin.H{
		"status":  "success",
		"message": msg,
		"name":    name,
	})
}

func parseName(s string) (string, bool) {
	switch s {
	case "ci":
		return name(s), true
	case "keke":
		return name(s), true
	case "cax":
		return name(s), true
	default:
		return s, false
	}
}

func name(name string) string {
	arr := utils.ParsedConf[name]
	msg := ""
	for _, sen := range arr {
		msg += sen + "！"
	}
	return msg
}
