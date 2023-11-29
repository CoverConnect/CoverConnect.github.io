package main

import (
	"encoding/base64"
	"syscall/js"
)

func main() {

	ch := make(chan struct{}, 0)
	funcs := js.Global().Get("Object").New()
	funcs.Set("decode", js.FuncOf(decode))
	funcs.Set("encode", js.FuncOf(encode))

	js.Global().Set("functions", funcs)
	<-ch
}

// java singature
func decode(this js.Value, args []js.Value) interface{} {
	rawDecodedText, err := base64.StdEncoding.DecodeString(args[0].String())
	if err != nil {
		return "decode error:" + err.Error()
	}
	return string(rawDecodedText)
}

func encode(this js.Value, args []js.Value) interface{} {
	str := args[0].String()
	encodeStr := base64.StdEncoding.EncodeToString([]byte(str))
	return string(encodeStr)
}
