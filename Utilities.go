package main

import (
	"encoding/json"
	"fmt"
)

func printIferr(err error) {
	if err != nil {
		print(err)
	}
}

func printStructObject(structObject interface{}) {
	jsonObject, _ := json.MarshalIndent(&structObject, "", "  ")
	fmt.Printf("MarshalIndent funnction output\n %s\n", string(jsonObject))
}
