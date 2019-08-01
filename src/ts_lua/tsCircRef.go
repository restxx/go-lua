// 提前GC并不会释放内存 之后的lua脚本仍可以正常运行
// 因为指針己经传入了luaState 除非先关闭luaState
// recuTest 相互引用无法释放

package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
	"unsafe"
)

const num = 1024 * 1024

var wg sync.WaitGroup

type Father struct {
	Data  [num]byte
	MySon *Son
}

func (f *Father) SetSon(s *Son) {
	f.MySon = s
}

type Son struct {
	Data  [num]byte
	MyDad *Father
}

func (s *Son) SetDad(f *Father) {
	s.MyDad = f
}

func test() {
	fmt.Println(debug.SetGCPercent(1))

	wg.Add(1)
	L := lua.NewState()
	// defer L.Close()

	f := &Father{}
	fmt.Println(unsafe.Sizeof(*f))
	L.SetGlobal("f", luar.New(L, f))

	runtime.SetFinalizer(f, func(_f *Father) {
		fmt.Println("Father内存回收")
	})

	s := &Son{}
	L.SetGlobal("s", luar.New(L, s))

	runtime.SetFinalizer(s, func(_s *Son) {
		fmt.Println("Son内存回收")
	})

	// L.Close()
	runtime.GC()
	// L.Close()

	if err := L.DoString(`
				f:SetSon(s)
				s:SetDad(f)
					 	`); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	wg.Done()
}

func recuTest() {
	fmt.Println(debug.SetGCPercent(1))

	f := &Father{}
	runtime.SetFinalizer(f, func(_f *Father) {
		fmt.Println("Father内存回收")
	})

	s := &Son{}
	runtime.SetFinalizer(s, func(_s *Son) {
		fmt.Println("Son内存回收")
	})

	f.MySon = s
	s.MyDad = f

	runtime.GC()
}

func main() {

	// test()
	// wg.Wait()

	recuTest()

	time.Sleep(6 * time.Second)
}
