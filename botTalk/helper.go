package botTalk

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func EntryToResponse(x Entry) Response {
	res := Response{}
	if x.meme != "" {
		res.memePath = x.meme
		res.msgType = MEME
	} else {
		res.text = x.words
		res.msgType = TEXT
	}
	return res
}

func responseToMessage(update tgbotapi.Update,response Response) tgbotapi.Chattable {
	switch response.msgType {
	case TEXT:
		return tgbotapi.NewMessage(update.Message.Chat.ID, response.text)
	case MEME:
		return tgbotapi.NewAnimationUpload(update.Message.Chat.ID, response.memePath)
	default:
		return tgbotapi.NewMessage(update.Message.Chat.ID, response.text)
	}

}

