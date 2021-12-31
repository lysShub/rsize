package main

import (
	"fmt"
	"unsafe"
)

func s() {
	/*
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
	*/
	var a []string = make([]string, 6, 10)
	a[0] = "aaaaaaa"

	ptrPtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 0))
	len := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 8))
	cap := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 16))

	fmt.Println(*len)
	fmt.Println(*cap)

	fmt.Println(ptrPtr)
}
