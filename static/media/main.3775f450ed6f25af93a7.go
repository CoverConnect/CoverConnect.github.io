package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"gopkg.in/yaml.v2"
)

func main() {

	ch := make(chan struct{}, 0)
	funcs := js.Global().Get("Object").New()
	funcs.Set("json2Yaml", js.FuncOf(json2Yaml))
	funcs.Set("yaml2Json", js.FuncOf(yaml2Json))

	js.Global().Set("functions", funcs)
	<-ch
}

// java singature
func json2Yaml(this js.Value, args []js.Value) interface{} {
	jsonData := args[0].String()
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		js.Global().Call("Error", err.Error())
		return ""
	}

	// Marshal the map into YAML
	yamlData, err := yaml.Marshal(&data)
	if err != nil {
		js.Global().Call("Error", err.Error())
		return ""
	}
	return string(yamlData)
}

func yaml2Json(this js.Value, args []js.Value) interface{} {
	// Unmarshal the YAML into a map
	yamlData := args[0].String()
	var data interface{}
	err := yaml.Unmarshal([]byte(yamlData), &data)
	if err != nil {
		js.Global().Call("Error", err.Error())
		return ""
	}
	data = convertMapI2MapS(data)

	// Marshal the map into JSON
	jsonData, err := json.Marshal(&data)
	if err != nil {
		js.Global().Call("Error", err.Error())
		return ""
	}

	return string(jsonData)
}
func convertMapI2MapS(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[fmt.Sprint(k)] = convertMapI2MapS(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convertMapI2MapS(v)
		}
	}
	return i
}