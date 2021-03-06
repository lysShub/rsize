package main

import (
	"fmt"
	"unsafe"
)

func array() {
	var a [4]int = [4]int{3, 3, 3, 3}

	/*
		type Array struct {
			data  unsafe.Pointer
			len int64
		}
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

func array2() {

	var a [2][2]int = [2][2]int{{1, 1}, {2, 2}}
	l := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 16))
	fmt.Println(*l) // 1
}
