package rsize

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Test() {
	// var a map[int]string = map[int]string{
	// 	1: "efsf",
	// 	2: "ewefs",
	// }

	var l string = "dsadfsaas"
	r := reflect.TypeOf(l)
	c := r.String()
	fmt.Println(c)

	var a map[int]int = map[int]int{
		1: 1,
		2: 2,
	}

	t := reflect.TypeOf(a)

	ts := t.String()

	b := (*maptype)(unsafe.Pointer(&a))
	fmt.Println(b, ts)
}

func eface(dataPtr, typePtr unsafe.Pointer) (size int) {

	efaceKind := *(*uint8)(unsafe.Pointer(uintptr(typePtr) + 2*word + 7))
	efaceKind = efaceKind & kindMask
	if tsize := originKind(efaceKind); tsize != 0 {
		size = tsize
	} else {
		switch efaceKind {
		case kindArray:
			// in this case, eface._type is arraytype
			size = earray(dataPtr, typePtr)
		case kindChan:
			size = echan(dataPtr, typePtr)
		case kindFunc:
		case kindInterface:
			size = eface(dataPtr, typePtr)
		case kindMap:
			size = emaps(dataPtr, typePtr)
		case kindPtr:
			size = eptrs(dataPtr, typePtr)
		case kindSlice:
			size = eslice(dataPtr, typePtr)
		case kindString:
			size = *(*int)(unsafe.Add(dataPtr, word))
		case kindStruct:
			size = estruct(dataPtr, typePtr)
		case kindUnsafePointer:
			size = eunptr(dataPtr, typePtr)
		default:
		}
	}
	return size
}

// getArrayEfaceElemType
//  get array element type when eface type is array
func getArrayEfaceElemType(efaceTypePtr *unsafe.Pointer) (typ uint8) {
	elemTypePtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(*efaceTypePtr) + word*4 + 16))
	eArrayElemKind := (*uint8)(unsafe.Pointer(uintptr(*elemTypePtr) + (2*word + 7)))

	return *eArrayElemKind
}

// size1 := (*uintptr)(unsafe.Pointer(uintptr(*etyp) + 0))
// prtdata := (*uintptr)(unsafe.Pointer(uintptr(*etyp) + word))
// hash := (*uint32)(unsafe.Pointer(uintptr(*etyp) + word*2))
// tflag := (*uint8)(unsafe.Pointer(uintptr(*etyp) + word*2 + 4))
// align := (*uint8)(unsafe.Pointer(uintptr(*etyp) + word*2 + 5))
// fieldAlign := (*uint8)(unsafe.Pointer(uintptr(*etyp) + word*2 + 6))
// fmt.Print(size1, prtdata, hash, tflag, align, fieldAlign)

// getSize 根据_type获取数据类型, t是_type的指针
//  如果是确定大小的数据, size将不为0
func getSize(vPtr unsafe.Pointer) (size int) {
	// 8 + 8 + 4 + 1 + 1 + 1
	kind := (*uint8)(unsafe.Pointer(uintptr(vPtr) + 23))

	if size := originKind(*kind); size != 0 {
		return size
	} else {
		switch *kind {
		case kindArray:
			eType := gotElemType(vPtr)
			counts := *(*uintptr)(unsafe.Pointer(uintptr(vPtr) + 0))
			if eSize := originKind(eType); eSize != 0 {
				return int(counts) * eSize
			} else {

			}
		}
	}
	return 0
}

// gotElemType 获取“数据”类型的元素的类型
//  Golang容器数据类型有：array chan slice ptr
// 	在src/runtime/type.go中定义有 arraytype等类似的结构体
// 	都有规律：第一个字段时_type, 第二个字段是*_type, 第二个字段就是元素的类型
func gotElemType(vPtr unsafe.Pointer) (elemType uint8) {
	eKindPtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(vPtr) + 48))
	elemType = *(*uint8)(unsafe.Pointer(uintptr(*eKindPtr) + 23))
	return elemType
}
