package main

import (
	"encoding/json"
	"fmt"
)

func replaceAt(index int, str string, token string) string {
	var updatedStr string
	for i, v := range str {
		updatedStr += string(v)
		if i == index {
			updatedStr += token
		}
	}
	return updatedStr
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
