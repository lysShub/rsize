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
	// var i interface{} = [2][2]string{{"9", "9"}, {"9", "9sfasqweras"}}
	// var i interface{} = [][]string{{"9", "9"}, {"9", "9sfasqweras"}}
	// var i interface{} = []string{"9", "9", "9", "9sfasqweras"}
	// var i interface{} = [4]string{"9", "9", "9", "9sfasqweras"}
	var i interface{} = [][]int{{1, 1, 1}, {1}}
	fmt.Println(rsize.GetEfaceSize(&i))
	return

	var a [4]int = [4]int{3, 3, 3, 3}

	/*
		 加如array的struct是：
			type Array struct {
				data  unsafe.Pointer
				len int64
			}
		 那么l应该是length

	*/
	l := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 8))
	fmt.Println(*l) // 3

	// data: *str1_data str1_len *str2_data str2_len
	var b [4]string = [4]string{"abc", "bbb", "c", "d"}

	// 获取第一个字符串
	s1 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 0))
	fmt.Println(*s1)
	// 获取第一个字符串长度长度
	ll := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 8))
	fmt.Println(*ll)

	// 取第二个字符串
	s2 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 16))
	fmt.Println(*s2)

	// 获取第1个字符串第n个字符
	data1Ptr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 0))
	// 第一个字符
	data1 := (*uint8)(unsafe.Pointer(uintptr(*data1Ptr) + 0))
	fmt.Println(string([]byte{*data1}))
	// 第二个字符
	data2 := (*uint8)(unsafe.Pointer(uintptr(*data1Ptr) + 1))
	fmt.Println(string([]byte{*data2}))

}

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
