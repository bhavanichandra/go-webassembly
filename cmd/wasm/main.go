package main

import (
	"encoding/json"
	"log"
	"syscall/js"
)

func main() {

	log.Println("Hello World from Webbrowser")
	js.Global().Set("JSONFormatter", jsonWrapper())
	<-make(chan int)
}
func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "       ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed!"
		}
		inputJson := args[0].String()
		log.Printf("input %s\n", inputJson)
		pretty, err := prettyJson(inputJson)
		if err != nil {
			log.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	})
}
