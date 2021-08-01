package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

type person struct{ name string }

//go:noinline
func (p *person) speak() {
	fmt.Println("my name is ", p.name)
}

func patchInstanceFunc() {
	p := &person{"Lance"}
	fmt.Println("original function output:")
	p.speak()

	patchGuard := sm.Patch((*person).speak, func(*person) {
		fmt.Println("we are all the same")
	})
	fmt.Println("after patch, function output:")
	p.speak()

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	p.speak()

	patchGuard.Restore()
	fmt.Println("restore, then output:")
	p.speak()

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	p.speak()

	fmt.Println()
}
