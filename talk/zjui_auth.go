package talk

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthZJUI(c *gin.Context) {
	q1 := c.DefaultPostForm("q1", "")
	q2 := c.DefaultPostForm("q2", "")

	res, ok := checkAnswers(q1, q2)

	c.JSON(200, gin.H{
		"ok": fmt.Sprintf("%t", ok),
		"message": res,
	})
}

func checkAnswers(q1, q2 string) (string, bool) {
	if q1 == "罗" && q2 == "逄" {
		return "Greetings to all aspiring people studying THE GOLD!", true
	}
	if q1 == "" || q2 == "" {
		return "请回答以上问题，出于一些原因我们需要核实您的身份", false
	}
	return "回答错误", false
}