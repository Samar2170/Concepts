package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)
}
func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

//	func jsonWrapper() js.Func {
//		jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
//			if len(args) != 1 {
//				return js.Null()
//			}
//			inputJson := args[0].String()
//			fmt.Printf("inputJson: %s", inputJson)
//			pretty, err := prettyJson(inputJson)
//			if err != nil {
//				return js.Null()
//			}
//			return pretty
//		})
//		return jsonFunc
//	}
func jsonWrapper() js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			result := map[string]any{
				"error": "invalid number of arguments",
			}
			return result
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "no document",
			}
			return result
		}
		jsonOpArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOpArea.Truthy() {
			result := map[string]any{
				"error": "no jsonoutput",
			}
			return result
		}
		inputJson := args[0].String()
		fmt.Printf("inputJson: %s", inputJson)
		pretty, err := prettyJson(inputJson)
		if err != nil {
			result := map[string]any{
				"error": err.Error(),
			}
			return result
		}
		jsonOpArea.Set("value", pretty)
		return nil
	})
	return jsonfunc
}
