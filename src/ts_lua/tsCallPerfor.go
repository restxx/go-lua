// 函数调用测试， 由luar导出调用10万次  luar导出逻辑消耗无明显差异 2.32秒
// 直接调用10万次 0.003秒

/*
lua 调用 go函数100000 次
2.3201327
go 调用 go Func 100000 次
0.0020001
lua 自加 100000次
inside lua 0.017001000000000044
0.017001
ret=  5000050000
go add 100000*1000 times
0.0280016
lua 调用SetToken 100000 times
2.1921254
*/

package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"log"
	"time"
)

func _tsFunc(str string, extra ...int) {
	switch str {
	case "a":
		if len(extra) != 3 || extra[0] != 1 || extra[1] != 2 || extra[2] != 3 {
			log.Fatalf("unexpected variable arguments: %v", extra)
		}
	case "b":
		if len(extra) != 0 {
			log.Fatalf("unexpected variable arguments: %v", extra)
		}
	case "c":
		if len(extra) != 1 || extra[0] != 4 {
			log.Fatalf("unexpected variable arguments: %v", extra)
		}
	}
}

func TESTLUA(count int, fn func()) {

	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("fn", luar.New(L, _tsFunc))

	start := time.Now()

	for i := 0; i <= count; i++ {

		// L.SetGlobal("fn", luar.New(L, _tsFunc))
		if err := L.DoString(`
		return fn("a", 1,2,3)
		`); err != nil {
			panic(err)
		}
	}

	fmt.Println(time.Now().Sub(start).Seconds())

}

func TESTGO(count int, fn func()) {
	start := time.Now()

	for i := 0; i <= count; i++ {

		_tsFunc("a", 1, 2, 3)
	}

	fmt.Println(time.Now().Sub(start).Seconds())
}

func luaAdd() {
	L := lua.NewState()
	defer L.Close()

	start := time.Now()

	if err := L.DoString(`
		function Add()
			local start = os.clock()
			j = 0
			for i=0, 100000, 1 do
				j = i+j
			end
			print("inside lua "..os.clock()-start)	
			return j
		end
		
		return Add()
		
		`); err != nil {
		panic(err)
	}
	fmt.Println(time.Now().Sub(start).Seconds())
	fmt.Println("ret= ", L.Get(-1))
}

func goAdd() int {
	start := time.Now()
	j := 0
	for i := 0; i <= 100000*1000; i++ {
		j += i
	}
	fmt.Println(time.Now().Sub(start).Seconds())
	return j
}

// ---------------------------------------------------------
type Use struct {
	Name  string
	token string
}

func (u *Use) SetToken(t string) {
	u.token = t
}

func (u *Use) Token() string {
	return u.token
}

func (u *Use) SetToken1(t string) {
	u.token = t
}

func (u *Use) Token1() string {
	return u.token
}

func (u *Use) SetToken2(t string) {
	u.token = t
}

func (u *Use) Token2() string {
	return u.token
}

func (u *Use) SetToken3(t string) {
	u.token = t
}

func (u *Use) Token3() string {
	return u.token
}

func (u *Use) SetToken4(t string) {
	u.token = t
}

func (u *Use) Token4() string {
	return u.token
}

func (u *Use) SetToken5(t string) {
	u.token = t
}

func (u *Use) Token5() string {
	return u.token
}

func tsLuarNew() {
	L := lua.NewState()
	defer L.Close()
	start := time.Now()
	for i := 1; i <= 100000; i++ {
		u := &Use{
			Name: "Tim",
		}
		L.SetGlobal("u", luar.New(L, u))

		if err := L.DoString(`
				u:SetToken("12345")
				`); err != nil {
			panic(err)
		}
	}
	fmt.Println(time.Now().Sub(start).Seconds())
}

func main() {

	fmt.Println("lua call go func 100000 times")
	TESTLUA(10000*10, nil)
	fmt.Println("go call go func 100000 times")
	TESTGO(10000*10, nil)
	fmt.Println("lua add 100000 times")
	luaAdd()
	fmt.Println("go add 100000*1000 times")
	goAdd()
	fmt.Println("lua 调用SetToken 100000 times")
	tsLuarNew()
}
