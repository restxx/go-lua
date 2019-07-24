package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

type Person struct {
	Name string
}

const luaPersonTypeName = "person"

// Registers my person type to given L.
func registerPersonType(L *lua.LState) {
	mt := L.NewTypeMetatable(luaPersonTypeName)
	L.SetGlobal("person", mt)
	// static attributes
	L.SetField(mt, "new", L.NewFunction(newPerson))
	// methods
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), personMethods))
}

// Constructor
func newPerson(L *lua.LState) int {
	person := &Person{L.CheckString(1)}
	ud := L.NewUserData()
	ud.Value = person
	L.SetMetatable(ud, L.GetTypeMetatable(luaPersonTypeName))
	L.Push(ud)
	return 1
}

// Checks whether the first lua argument is a *LUserData with *Person and returns this *Person.
func checkPerson(L *lua.LState) *Person {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Person); ok {
		return v
	}
	L.ArgError(1, "person expected")
	return nil
}

var personMethods = map[string]lua.LGFunction{
	"name": personGetSetName,
}

// Getter and setter for the Person#Name
func personGetSetName(L *lua.LState) int {

	fmt.Println(L.GetTop())
	// 取出Person
	p := checkPerson(L)
	// 用栈的情况在判断是get还是set
	if L.GetTop() >= 2 { // setName
		p.Name = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(p.Name)) // getName
	return 1
}

func main() {
	L := lua.NewState()
	defer L.Close()
	registerPersonType(L)
	if err := L.DoFile("D:\\GO_SOURCE\\go_lua\\src\\ts_lua\\person.lua"); err != nil {
		panic(err)
	}
}
