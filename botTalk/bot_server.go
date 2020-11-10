package botTalk

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func (x *JinTalkBotServer) Init() {
	x.dictFile, x.token = commandLineInit()

	var err error
	x.bot, err = tgbotapi.NewBotAPI(x.token)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Authorized on account %s", x.bot.Self.UserName)

	x.bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	x.updates, err = x.bot.GetUpdatesChan(u)
}

func (x *JinTalkBotServer) Configure() {
	x.totalFreq = 0
	dat, err := ioutil.ReadFile(x.dictFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("==== config ===\n")
	lines := strings.Split(string(dat), "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fmt.Printf("%s\n", line)
		buildLine(x, line)
	}
	x.randomSrc = rand.NewSource(time.Now().UnixNano())
	fmt.Printf("==== config end ===\n")
	fmt.Printf("init success\n")
	fmt.Printf("total normal response frequency: %d\n\n", x.totalFreq)
}

func (x *JinTalkBotServer) Start() {
	for update := range x.updates {
		// ignore any non-Message Updates
		if update.Message == nil {
			continue
		}

		response := botSay(x)
		msg := responseToMessage(update, response)

		_, err := x.bot.Send(msg)
		if err != nil {
			fmt.Printf("err = %v", err)
		}
	}
}

