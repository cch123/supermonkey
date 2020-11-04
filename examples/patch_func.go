package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

func patchFunc() {
	fmt.Println("original function output:")
	heyHey()

	patchGuard := sm.Patch(heyHey, func() {
		fmt.Println("please be polite")
	})
	fmt.Println("after patch, function output:")
	heyHey()

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	heyHey()

	fmt.Println()
}

//go:noinline
func heyHey() {
	fmt.Println("fake")
}
