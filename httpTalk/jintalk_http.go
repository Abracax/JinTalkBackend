package httpTalk

import (
	"github.com/Abracax/JinTalkBackend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

type JinTalkHttpServer struct {
	router *gin.Engine
}

func (s *JinTalkHttpServer) Init() {
	s.router = gin.Default()
	if err := utils.ParseConf(); err != nil {
		log.Printf("init conf err=%v", err)
		return
	}
}

func (s *JinTalkHttpServer) Configure() {
	rand.Seed(time.Now().UTC().UnixNano())
	s.router.GET("/httpTalk/:name", Talk)
	s.router.GET("/httpTalk", Talk)
	s.router.GET("/random", Random)
	s.router.POST("/auth", AuthZJUI)
}

func (s *JinTalkHttpServer) Start() {
	err := s.router.Run(":2333")
	if err != nil {
		panic("run router err")
	}
}