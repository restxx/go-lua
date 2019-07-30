// gopher-luar测试 golang结构指針到lua
// go函数到lua
package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
	"log"
)

type User struct {
	Name  string
	token string
}

func (u *User) SetToken(t string) {
	u.token = t
}

func (u *User) Token() string {
	return u.token
}

const script = `
print("Hello from Lua, " .. u.Name .. "!")
u:SetToken("12345")
`

func tsStruct() {
	L := lua.NewState()
	defer L.Close()

	u := &User{
		Name: "Tim",
	}
	L.SetGlobal("u", luar.New(L, u))
	if err := L.DoString(script); err != nil {
		panic(err)
	}

	fmt.Println("Lua set your token to:", u.Token())
	// Output:
	// Hello from Lua, Tim!
	// Lua set your token to: 12345
}

func tsFunc2lua() {
	L := lua.NewState()
	defer L.Close()

	fn := func(str string, extra ...int) {
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

	L.SetGlobal("fn", luar.New(L, fn))
	if err := L.DoString(`
		return fn("a", 1,2,3)
		`); err != nil {
		panic(err)
	}
}

func main() {
	// tsStruct()
	tsFunc2lua()
}
