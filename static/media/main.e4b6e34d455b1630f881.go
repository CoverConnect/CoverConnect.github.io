package main

import (
	"syscall/js"
	"time"
)

func main() {

	ch := make(chan struct{}, 0)
	funcs := js.Global().Get("Object").New()
	funcs.Set("convertTimestamp", js.FuncOf(convertTimestamp))

	js.Global().Set("functions", funcs)
	<-ch
}

func convertTimestamp(this js.Value, p []js.Value) interface{} {
	if len(p) != 1 {
		return "Invalid number of arguments. Expected 1."
	}

	timestamp := p[0].Int()
	convertedTime := time.Unix(int64(timestamp), 0)

	result := map[string]interface{}{
		"year":   convertedTime.Year(),
		"month":  int(convertedTime.Month()),
		"day":    convertedTime.Day(),
		"hour":   convertedTime.Hour(),
		"minute": convertedTime.Minute(),
		"second": convertedTime.Second(),
	}

	return js.ValueOf(result)
}
