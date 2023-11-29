package main

import (
    "bytes"
    "encoding/base64"
    "image"
    _ "image/jpeg"
    "image/png"
    "strings"
    "syscall/js"

    "github.com/tuotoo/qrcode"
)
func decodeQR(this js.Value, args []js.Value) interface{} {
    imageData := args[0].String()

    // Remove the data URL prefix and decode the base64 string
    reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(strings.Split(imageData, ",")[1]))

    // Decode the image
    img, _, err := image.Decode(reader)
    if err != nil {
        return err.Error()
    }

    // Decode the QR code
    qrCode, err := qrcode.Decode(bytes.NewReader(imgToByte(img)))
    if err != nil {
        return err.Error()
    }

    return js.ValueOf(qrCode.Content)
}

func imgToByte(img image.Image) []byte {
    var buf bytes.Buffer
    _ = png.Encode(&buf, img) // Encode as PNG
    return buf.Bytes()
}

func main() {

	ch := make(chan struct{}, 0)
	funcs := js.Global().Get("Object").New()
	funcs.Set("decodeQR", js.FuncOf(decodeQR))

	js.Global().Set("functions", funcs)
	<-ch
}