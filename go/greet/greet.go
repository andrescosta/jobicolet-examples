package main

import (
	"github.com/andrescosta/jobicolet-sdk-go/pkg/sdk"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export init
func _init() {
	sdk.OnEvent = test
}

func test(id uint32, data string) (uint64, string) {
	sdk.Log(id, sdk.InfoLevel, "Hello from Tiny GO")
	sdk.Log(id, sdk.DebugLevel, data)
	sdk.Log(id, sdk.InfoLevel, data)
	return sdk.NoError, "ok"
}
