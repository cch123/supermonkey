package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

func patchFuncSymbol() {
	fmt.Println("original function output:")
	heyHeyHey()

	patchGuard := sm.PatchByFullSymbolName("main.heyHeyHey", func() {
		fmt.Println("please be polite")
	})
	fmt.Println("after patch, function output:")
	heyHeyHey()

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	heyHeyHey()

	fmt.Println()
}

//go:noinline
func heyHeyHey() {
	fmt.Println("fake")
}
