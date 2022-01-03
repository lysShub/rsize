package rsize

import (
	"unsafe"
)

func GetEfaceSize(efacePtr *interface{}) (size int) {
	efaceTypePtr := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(efacePtr)) + word*0))
	efaceDataPtr := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(efacePtr)) + word*1))

	return eface(efaceDataPtr, efaceTypePtr)
}

func eface(dataPtr, typePtr unsafe.Pointer) (size int) {
	efaceKind := (*uint8)(unsafe.Pointer(uintptr(typePtr) + 2*word + 7))
	if size = originKind(*efaceKind); size != 0 {
		return
	} else {
		switch *efaceKind {
		case kindArray:
			// in this case, eface._type is arraytype
			return earray(dataPtr, typePtr)
		case kindChan:
		case kindFunc:
		case kindInterface:
		case kindMap:
		case kindPtr:
		case kindSlice:
			return eslice(dataPtr, typePtr)
		case kindString:
		case kindStruct:
		case kindUnsafePointer:
		default:
			return 0
		}
	}
	return 0
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

// originKind
// 	如果类型是纯纯的“值类型”, 就返回此类型的大小, 否则返回0
func originKind(kind uint8) (size int) {
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
