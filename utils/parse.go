package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var ParsedConf = make(map[string][]string)

func ParseConf() error {
	f, err := ioutil.ReadFile("./conf/jintalk.json")
	if err != nil {
		log.Printf("read conf err=%v", err)
		return err
	}

	if err = json.Unmarshal(f, &ParsedConf); err != nil {
		log.Printf("conf parse to json err=%v", err)
		return err
	}

	log.Printf("test=%s", ParsedConf["greetings"][0])
	return nil
}
