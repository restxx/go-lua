// 测试10000个goroutine调用luastate内存占用
package main

import (
	lua "github.com/yuin/gopher-lua"
	"time"
)

// 内存占用1800M 单个180K
func test1() {
	for i := 0; i < 10000; i++ {
		L := lua.NewState()
		defer L.Close()
		if err := L.DoString(`
			function fib(n)
				if n < 2 then
					return 1
				end
				return fib(n-1) + fib(n-2)
			end
			print(fib(10))
		`); err != nil {
			panic(err)
		}
	}
	time.Sleep(100 * time.Second)
}

func _ts() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`
			function fib(n)
				if n < 2 then
					return 1
				end
				return fib(n-1) + fib(n-2)
			end
			print(fib(10))
			`); err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
}

// 占用内存2135MB 単个210K
func test2() {
	for i := 0; i < 10000; i++ {
		go _ts()
	}
	time.Sleep(20 * time.Second)
}

func main() {
	// test1()
	test2()
}
