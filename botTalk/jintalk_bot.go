package botTalk

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func buildLine(x *JinTalkBotServer, line string) {
	var ent Entry
	var followFreq int
	tokens := strings.Split(line, " ")
	n := len(tokens)
	for i := 0; i < n; {
		// say [xxx] and meme
		if tokens[i] == "say" {
			next := tokens[i+1]
			if next == "meme" {
				ent.meme = tokens[i+2]
				i = i + 3
			} else {
				ent.words = next
				i = i + 2
			}
			continue
		}
		// with freq [n]
		if tokens[i] == "with" {
			var err error
			ent.freq, err = strconv.Atoi(tokens[i+2])
			if err != nil {
				panic("bad configurations")
			}
			i = i + 3
			continue
		}
		// follows freq [m]
		if tokens[i] == "follows" {
			var err error
			followFreq, err = strconv.Atoi(tokens[i+2])
			if err != nil {
				panic("bad configurations")
			}
			i = i + 3
			continue
		}
	}
	if len(x.entries) != 0 {
		x.entries[len(x.entries)-1].followFreq = followFreq
	}
	x.entries = append(x.entries, ent)
	x.totalFreq = x.totalFreq + ent.freq
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

func Say(x *JinTalkBotServer) Response {
	r := rand.New(x.randomSrc)
	var index int
	if x.lastIndex != -1 {
		index = r.Intn(x.totalFreq + x.entries[x.lastIndex].followFreq)
	} else {
		index = r.Intn(x.totalFreq)
	}
	if index >= x.totalFreq {
		x.lastIndex++
		return EntryToResponse(x.entries[x.lastIndex])
	}
	entryN := 0
	innerN := x.entries[0].freq
	for i := 0; i < index; i++ {
		innerN--
		if innerN == 0 {
			entryN++
			innerN = x.entries[entryN].freq
		}
	}
	x.lastIndex = entryN
	return EntryToResponse(x.entries[entryN])
}

func (x *JinTalkBotServer) Start() {
	for update := range x.updates {
		// ignore any non-Message Updates
		if update.Message == nil {
			continue
		}

		response := Say(x)
		msg := responseToMessage(update, response)

		_, err := x.bot.Send(msg)
		if err != nil {
			fmt.Printf("err = %v", err)
		}
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
