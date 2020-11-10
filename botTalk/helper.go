package botTalk

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

func entryToResponse(x Entry) Response {
	res := Response{}
	if x.meme != "" {
		res.memePath = x.meme
		res.msgType = MEME
	} else if len(x.seq) > 0 {
		res.msgType = SEQ
		res.seq = x.seq
	} else {
		res.text = x.words
		res.msgType = TEXT
	}
	return res
}

func responseToMessage(update tgbotapi.Update,response Response) (tgbotapi.Chattable, bool) {
	switch response.msgType {
	case TEXT:
		return tgbotapi.NewMessage(update.Message.Chat.ID, response.text), true
	case MEME:
		return tgbotapi.NewAnimationUpload(update.Message.Chat.ID, response.memePath), true
	case SEQ:
		return nil, false
	default:
		return nil, true
	}
}

func commandLineInit() (string, string) {
	argv := os.Args
	argc := len(argv)
	if argc != 3 {
		panic("args mismatch")
	}
	fmt.Printf("using dict file: [%s]\n", argv[1])

	return argv[1], argv[2]
}
