package main

import lua "github.com/yuin/gopher-lua"

// 将ubyte传入lua并打印

type TClass struct {
	Buf []byte
}

func NewTClass(buf []byte) *TClass {

	return &TClass{Buf: buf}
}

func LuaNewTClass(L *lua.LState) int {

	return 0
}
