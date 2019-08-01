package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"time"
)

func main() {
	start := time.Now()
	var ret int
	for i := 0; i < 30; i++ {
		ret = GoFib(i)
	}
	fmt.Println(time.Now().Sub(start).Seconds())
	fmt.Println(ret)

	LuaFib(30)
}

func GoFib(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return GoFib(n-2) + GoFib(n-1)
}

func LuaFib(n int) {
	L2 := lua.NewState()
	defer L2.Close()

	if err := L2.DoFile("D:\\GO_SOURCE\\go_lua\\src\\ts_lua\\lua_script\\Fib.lua"); err != nil {
		panic(err)
	}

	// go 调用lua中的max函数
	fn := L2.GetGlobal("tsFib")
	P := lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true}

	_ = L2.CallByParam(P,
		lua.LNumber(n),
	)

	fmt.Println(L2.Get(-1))
}
