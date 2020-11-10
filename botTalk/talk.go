package botTalk

import (
	"math/rand"
	"strconv"
	"strings"
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

func botSay(x *JinTalkBotServer) Response {
	r := rand.New(x.randomSrc)
	var index int
	if x.lastIndex != -1 {
		index = r.Intn(x.totalFreq + x.entries[x.lastIndex].followFreq)
	} else {
		index = r.Intn(x.totalFreq)
	}
	if index >= x.totalFreq {
		x.lastIndex++
		return entryToResponse(x.entries[x.lastIndex])
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
	return entryToResponse(x.entries[entryN])
}
