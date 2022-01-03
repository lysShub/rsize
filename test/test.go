package main

import (
	"fmt"
	"unsafe"

	"github.com/lysShub/rsize"
)

func main() {
	// testSlice()
	// return

	// var ac int = 0 (unsafe.Pointer)(unsafe.Pointer(&ac))
	var i interface{} = [2][2]string{{"9", "9"}, {"9", "9sfasqweras"}}
	// var i interface{} = [][]string{{"9", "9"}, {"9", "9sfasqweras"}}
	// var i interface{} = []string{"9", "9", "9", "9sfasqweras"}
	// var i interface{} = [4]string{"9", "9", "9", "9sfasqweras"}
	// var i interface{} = []string{"1", "1", "1", "a1"}
	// var i interface{} = [][][]int{{{1, 3, 1}, {2}}, {{1}}}
	// var i interface{} = [][]int{{1, 2, 3}, {1}}
	// var i interface{} = [][]string{{"9", "9"}, {"9", "9sfasqweras"}}
	// var i interface{} = [][3]string{{"9", "9", "9"}, {"9", "9", "9sfasqweras"}}

	fmt.Println(rsize.GetEfaceSize(&i))

	return
	tp := *(*unsafe.Pointer)(unsafe.Pointer(&i))
	dp := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + 8))
	a := (*slicetype)(unsafe.Pointer(tp))
	fmt.Println(a)

	b := (*slice)(unsafe.Pointer(dp))

	sdp := (*unsafe.Pointer)(unsafe.Pointer(uintptr(dp) + 0))
	elp1 := (*slice1)(unsafe.Pointer(uintptr(*sdp) + 0))
	elp2 := (*slice1)(unsafe.Pointer(uintptr(*sdp) + 8*3))

	fmt.Println(b, elp1, elp2)
	return

}

type slice struct {
	array *tslice1 //unsafe.Pointer
	len   int
	cap   int
}

type tslice1 struct {
	sub1 slice1
	sub2 slice1
}

type slice1 struct {
	array unsafe.Pointer
	len   int
	cap   int
}

type slicetype struct {
	typ  _type
	elem *slicetype
}

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	equal      func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata     *byte
	str        nameOff
	ptrToThis  typeOff
}

type tflag uint8
type nameOff int32
type typeOff int32

func testSlice() {
	/*
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
	*/
	// var i interface{} = [][]string{
	// 	{"a", "b"},
	// 	{"c", "d"},
	// }
	// var i interface{} = []string{"a", "b"}
	var i interface{} = []int{1, 2, 3}

	sliceTypePtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + 0))
	sliceElemTypePtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(*sliceTypePtr) + 8*4 + 16))

	size := (*uintptr)(unsafe.Pointer(uintptr(*sliceTypePtr) + 0))
	kind := (*uint8)(unsafe.Pointer(uintptr(*sliceTypePtr) + 23))
	eKind := (*uint8)(unsafe.Pointer(uintptr(*sliceElemTypePtr) + 23))

	fmt.Println(*size, *kind, *eKind)
}
