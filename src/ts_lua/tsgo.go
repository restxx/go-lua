package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

// go 调用lua函数 lua调用go函数示例

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`print("hello")`); err != nil {
		panic(err)
	}

	L2 := lua.NewState()
	defer L2.Close()
	// 在lua中调用go函数 Double
	L2.SetGlobal("double", L2.NewFunction(Double)) /* Original lua_setglobal uses stack... */

	if err := L2.DoFile("D:\\GO_SOURCE\\go_lua\\src\\ts_lua\\lua_script\\hello.lua"); err != nil {
		panic(err)
	}

	// go 调用lua中的max函数
	fn := L2.GetGlobal("max")
	_ = L2.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true},
		lua.LNumber(3),
		lua.LNumber(5),
	)
	fmt.Println(L2.Get(-1))

	// test()
}

func Double(L *lua.LState) int {
	lv := L.ToInt(1)            /* get argument */
	L.Push(lua.LNumber(lv * 2)) /* push result */
	return 1                    /* number of results */
}
