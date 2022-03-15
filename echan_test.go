package rsize

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestChan(t *testing.T) {

	var a chan uint8 = make(chan uint8, 3)

	var b interface{} = a

	ptr := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 8))

	fmt.Println(ptr)

	// ptr := unsafe.Pointer(&a)

	echan(ptr, unsafe.Pointer(uintptr(0)))
}
