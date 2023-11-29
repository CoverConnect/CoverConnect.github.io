package main

import (
	"encoding/json"
	"strings"
	"syscall/js"
)

type Input struct {
	Text      string `json:"text"`
	Unescape  int    `json:"unescape"`
	Indent    int    `json:"indent"`
}

func main() {

	ch := make(chan struct{}, 0)
	funcs := js.Global().Get("Object").New()
	funcs.Set("formatJSON", js.FuncOf(formatJSON))

	js.Global().Set("functions", funcs)
	<-ch
}

// java singature
func formatJSON(this js.Value, args []js.Value) interface{} {
	inputJS := args[0].String()
	var input Input
	json.Unmarshal([]byte(inputJS), &input)

	var output interface{}
	json.Unmarshal([]byte(input.Text), &output)

	indentString := strings.Repeat(" ", input.Indent)
	formatted, err := json.MarshalIndent(output, "", indentString)
	if err != nil {
		return err.Error()
	}

	formattedString := string(formatted)
	for i := 0; i < input.Unescape; i++ {
		formattedString = strings.ReplaceAll(formattedString, "\\\"", "\"")
	}
	return formattedString
}