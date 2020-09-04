package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

func main() {
	fmt.Println("original function output:")
	heyHey()
	fmt.Println()

	sm.PatchByFullSymbolName("main.heyHey", func() {
		fmt.Println("please be polite")
	})
	fmt.Println("after patch, function output:")
	heyHey()
	fmt.Println()

	sm.UnpatchAll()
	fmt.Println("unpatch all, then output:")
	heyHey()
}

//go:noinline
func heyHey() {
	fmt.Println("fuck")
}
