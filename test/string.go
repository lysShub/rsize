package main

import (
	"fmt"
	"unsafe"
)

func str() {
	var s string = "1234sdfasdfas"

	l := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + 8))
	fmt.Println(*l)

	dataPtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + 0))
	data1 := (*uint8)(unsafe.Pointer(uintptr(*dataPtr) + 0))
	fmt.Println(string([]byte{*data1}))
}
