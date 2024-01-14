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
	sdk.Log(id, sdk.NoLevel, data+"_nolevel")
	sdk.Log(id, sdk.InfoLevel, data+"_info")
	sdk.Log(id, sdk.DebugLevel, data+"_debug")
	sdk.Log(id, sdk.WarnLevel, data+"_warn")
	sdk.Log(id, sdk.ErrorLevel, data+"_error")
	sdk.Log(id, sdk.FatalLevel, data+"_fatal")
	sdk.Log(id, sdk.PanicLevel, data+"_panic")
	return sdk.NoError, "ok"
}
