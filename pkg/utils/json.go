package utils

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
)

func JsonAssertion(src, dest interface{}) (err error) {
	var bts []byte
	if bts, err = json.Marshal(src); err != nil {
		return
	}

	if err = json.Unmarshal(bts, dest); err != nil {
		return
	}

	return
}

func PrintJson(input any) {
	b, err := json.MarshalIndent(input, "", "	")
	if err != nil {
		log.Error(err)
	}

	fmt.Println(string(b))
}
