package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

type person struct{ name string }

func (p *person) speak() {
	fmt.Println("my name is ", p.name)
}

func main() {
	var p = person{"Xargin"}
	fmt.Println("original function output:")
	p.speak()
	fmt.Println()

	sm.Patch("main", "*person", "speak", func() {
		fmt.Println("we are all the same")
	})
	fmt.Println("after patch, function output:")
	p.speak()
	fmt.Println()

	sm.UnpatchAll()
	fmt.Println("unpatch all, then output:")
	p.speak()
}
