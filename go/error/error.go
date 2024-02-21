package main

import (
	"github.com/andrescosta/jobicolet-sdk-go/pkg/sdk"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export init
func _init() {
	sdk.OnEvent = doerror
}

func doerror(data string) (uint64, string) {
	sdk.Log(sdk.ErrorLevel, data)
	return 500, data
}
