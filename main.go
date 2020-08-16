package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"playground/jin/talk"
	"playground/jin/utils"
	"time"
)

func main() {
	if err := utils.ParseConf(); err != nil {
		log.Printf("init conf err=%v", err)
		return
	}
	rand.Seed(time.Now().UTC().UnixNano())
	r := gin.Default()
	r.GET("/talk/:name", talk.Talk)
	r.GET("/talk", talk.Talk)
	r.GET("/random", talk.Random)
	//金老师的经典发言
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
