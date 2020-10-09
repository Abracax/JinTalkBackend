package main

import (
	"github.com/Abracax/JinTalkBackend/talk"
	"github.com/Abracax/JinTalkBackend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

func main() {
	if err := utils.ParseConf(); err != nil {
		log.Printf("init conf err=%v", err)
		return
	}
	rand.Seed(time.Now().UTC().UnixNano())
	router := gin.Default()
	router.GET("/talk/:name", talk.Talk)
	router.GET("/talk", talk.Talk)
	router.GET("/random", talk.Random)
	router.POST("/auth", talk.AuthZJUI)
	err := router.Run(":2333")
	if err != nil {
		log.Printf("run router err=%v", err)
		return
	}
}
