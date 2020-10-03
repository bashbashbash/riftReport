package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func replaceAt(index int, str string, token string) string {
	var tokenized = strings.Split(str, "***")
	var updated = ""
	for i, v := range tokenized {
		if i == index {
			updated += token
		} else {
			updated += v
		}
	}
	return updated
}

func printIferr(err error) {
	if err != nil {
		print(err)
	}
}

func printStructObject(structObject interface{}) {
	jsonObject, _ := json.MarshalIndent(&structObject, "", "  ")
	fmt.Printf("MarshalIndent funnction output\n %s\n", string(jsonObject))
}
