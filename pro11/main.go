package main

import "fmt"

type Man struct {
	name string
	age  int
}

func (m *Man) say() {
	fmt.Println("hello,I am " + m.name)
}

func (m *Man) setName1(name string) {
	m.name = name
}

func (m Man) setName2(name string) {
	m.name = name
}

func main() {
	m := Man{"a", 10}
	m.say()
	m1 := Man{"b", 20}
	m1.say()

	m.setName1("aa")
	m.say()

	m.setName2("bb")
	m.say()

}
