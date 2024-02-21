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

func test(data string) (uint64, string) {
	sdk.Log(sdk.NoLevel, data+"_nolevel")
	sdk.Log(sdk.InfoLevel, data+"_info")
	sdk.Log(sdk.DebugLevel, data+"_debug")
	sdk.Log(sdk.WarnLevel, data+"_warn")
	sdk.Log(sdk.ErrorLevel, data+"_error")
	sdk.Log(sdk.FatalLevel, data+"_fatal")
	sdk.Log(sdk.PanicLevel, data+"_panic")
	return sdk.NoError, "ok"
}
