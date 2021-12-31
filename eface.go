package rsize

import "unsafe"

func GetEfaceSize(e *interface{}) (size int) {
	v := (*unsafe.Pointer)(unsafe.Pointer(e))
	return getSize(v)

	// if size := originValue(v); size != 0 {
	// 	return size
	// } else {
	// 	switch kind {
	// 	case 17: // array

	// 		v := (*myarraytype)(*v)
	// 		fmt.Println(v)
	// 		return 0
	// 	case 99:
	// 	}
	// }

	// return 0
}

// originValue
// 	如果类型是纯纯的“值类型”, 就返回此类型的大小, 否则返回0
func originValue(kind uint8) (size int) {
	switch kind {
	case kindBool, kindInt8, kindUint8:
		return 1
	case kindInt16, kindUint16:
		return 2
	case kindInt32, kindUint32, kindFloat32:
		return 4
	case kindInt64, kindUint64, kindFloat64, kindComplex64, kindFunc:
		return 8
	case kindComplex128:
		return 16
	case kindInt, kindUint:
		return int(word)
	default:
		return 0
	}
}

// getSize 根据_type获取数据类型, t是_type的指针
//  如果是确定大小的数据, size将不为0
func getSize(vPtr *unsafe.Pointer) (size int) {
	// 8 + 8 + 4 + 1 + 1 + 1
	kind := (*uint8)(unsafe.Pointer(uintptr(*vPtr) + 23))

	if size := originValue(*kind); size != 0 {
		return size
	} else {
		switch *kind {
		case kindArray:
			eType := gotElemType(vPtr)
			counts := *(*uintptr)(unsafe.Pointer(uintptr(*vPtr) + 0))
			if eSize := originValue(eType); eSize != 0 {
				return int(counts) * eSize
			} else {

			}
		}
	}
	return 0
}

func array(vPtr *unsafe.Pointer) (size int) {

	return
}

// gotElemType 获取“数据”类型的元素的类型
//  Golang容器数据类型有：array chan slice ptr
// 	在src/runtime/type.go中定义有 arraytype等类似的结构体
// 	都有规律：第一个字段时_type, 第二个字段是*_type, 第二个字段就是元素的类型
func gotElemType(vPtr *unsafe.Pointer) (elemType uint8) {
	eKindPtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(*vPtr) + 48))
	elemType = *(*uint8)(unsafe.Pointer(uintptr(*eKindPtr) + 23))
	return elemType
}
