package main

import (
    "github.com/ishiikurisu/boinga"
    "fmt"
    "syscall/js"
)

func main() {
    js.Global().Set("toBoinga", js.FuncOf(toBoinga))
    js.Global().Set("fromBoinga", js.FuncOf(fromBoinga))
    fmt.Println("Go WASM is ready")
    <-make(chan bool)
}

func toBoinga(this js.Value, params []js.Value) interface{} {
    return js.ValueOf(boinga.ToBoinga(params[0].String()))
}

func fromBoinga(this js.Value, params []js.Value) interface{} {
    return js.ValueOf(boinga.FromBoinga(params[0].String()))
}
