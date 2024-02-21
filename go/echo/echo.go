package main

import (
	"github.com/andrescosta/jobicolet-sdk-go/pkg/sdk"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export init
func _init() {
	sdk.OnEvent = echo
}

func echo(data string) (uint64, string) {
	sdk.Log(sdk.NoLevel, data+"_nolevel")
	return 0, data
}
