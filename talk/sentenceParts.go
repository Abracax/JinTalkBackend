package talk

import (
	"math/rand"
	"playground/jin/utils"
)

func ParseName(s string) (string, bool) {
	switch s {
	case "ci":
		return name(s), true
	case "keke":
		return name(s), true
	case "cax":
		return name(s), true
	default:
		return s, false
	}
}

func name(name string) string {
	arr := utils.ParsedConf[name]
	msg := ""
	for _, sen := range arr {
		msg += sen + "！"
	}
	return msg
}

func GetRandFromConfEntry(s string) string {
	if confArr, ok := utils.ParsedConf[s]; ok {
		r := rand.Intn(len(confArr))
		msg := confArr[r]
		return msg
	}
	return ""
}
