package main

import (
	"github.com/Abracax/JinTalkBackend/botTalk"
	"github.com/Abracax/JinTalkBackend/httpTalk"
)

type JinTalkServer interface {
	Init()
	Configure()
	Start()
}

func main() {
	var httpServer httpTalk.JinTalkHttpServer
	var botServer botTalk.JinTalkBotServer
	StartServer(&botServer)
	StartServer(&httpServer)
}

func StartServer(s JinTalkServer) {
	s.Init()
	s.Configure()
	s.Start()
}