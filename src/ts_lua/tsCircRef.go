package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"runtime"
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
	wg.Add(1)
	L := lua.NewState()
	defer L.Close()

	f := &Father{Data: [num]byte{0}}
	fmt.Println(unsafe.Sizeof(*f))
	L.SetGlobal("f", luar.New(L, f))

	s := &Son{Data: [num]byte{0}}
	L.SetGlobal("s", luar.New(L, s))

	if err := L.DoString(`
				f:SetSon(s)
				s:SetDad(f)
					 	`); err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	wg.Done()
}

func main() {
	for i := 0; i < 5000; i++ {
		go test()
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	runtime.GC()
	time.Sleep(300 * time.Second)
}
