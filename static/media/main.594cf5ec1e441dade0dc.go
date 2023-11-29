package main

import (
	"encoding/base64"
	"syscall/js"

	"github.com/skip2/go-qrcode"
)

func main() {

	ch := make(chan struct{}, 0)
	funcs := js.Global().Get("Object").New()
	funcs.Set("generateQR", js.FuncOf(generateQR))

	js.Global().Set("functions", funcs)
	<-ch
}

func generateQR(this js.Value, args []js.Value) interface{} {
    content := args[0].String()
    qr, err := qrcode.Encode(content, qrcode.Medium, 256)
    if err != nil {
        return err.Error()
    }
    return base64.StdEncoding.EncodeToString(qr)
	// return content
}