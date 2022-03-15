package rsize

import (
	"fmt"
	"unsafe"
)

type ptrtype struct {
	typ  _type
	elem *_type
}

func eptrs(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {

	elemDataPtr := unsafe.Add(dataPtr, 0)
	elemTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))

	elemKind := (*(*uint8)(unsafe.Add(elemTypePtr, 2*word+7))) & kindMask

	if tsize := originKind(elemKind); tsize != 0 {
		size = tsize
	} else {
		switch elemKind {
		case kindArray:
			size = earray(elemDataPtr, elemTypePtr)
		case kindChan:
			size = echan(elemDataPtr, elemTypePtr)
		case kindFunc:
		case kindInterface:
			size = eface(elemDataPtr, elemTypePtr)
		case kindMap:
			size = emaps(elemDataPtr, elemTypePtr)
		case kindPtr:
			size = eptrs(elemDataPtr, elemTypePtr)
		case kindSlice:
			size = eslice(elemDataPtr, elemTypePtr)
		case kindString:
			size = size + *(*int)(unsafe.Add(dataPtr, word))
		case kindStruct:
			size = eslice(elemDataPtr, elemTypePtr)
		case kindUnsafePointer:
			size = eunptr(elemDataPtr, elemTypePtr)
		default:
		}
	}

	a := (*ptrtype)(typePtr)
	fmt.Println(a)

	b := *(*int)(dataPtr)
	fmt.Println(b, elemDataPtr, elemTypePtr)
	return size
}
