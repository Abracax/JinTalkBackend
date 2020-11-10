package botTalk

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
)

const (
	DICT string = "conf/jintalk.dict"
	TOKEN string = ""
)

type JinTalkBotServer struct {
	dictFile string
	token string

	updates tgbotapi.UpdatesChannel
	bot *tgbotapi.BotAPI

	totalFreq int
	entries   []Entry
	randomSrc rand.Source
	lastIndex int
}

type Entry struct {
	words      string
	freq       int
	followFreq int
	meme       string
}

type MsgType int32

type Response struct {
	msgType  MsgType
	text     string
	memePath string
}

const (
	TEXT MsgType = iota
	MEME
)

