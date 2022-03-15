package rsize

import (
	"unsafe"
)

func eslice(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {

	subElemTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))
	subElemKind := (*(*uint8)(unsafe.Add(subElemTypePtr, 2*word+7))) & kindMask

	sliceCounts := *(*int)(unsafe.Add(dataPtr, word*1)) // calculate len rather than cap
	if sliceCounts <= 0 {
		return 0
	}

	if size = originKind(subElemKind); size != 0 {
		size = int(sliceCounts) * size
	} else {
		// must ergodic all elements
		switch subElemKind {
		case kindArray:
			for i := 0; i < sliceCounts; i++ {
				size = size + earray(dataPtr, subElemTypePtr)
			}
		case kindChan:
			for i := 0; i < sliceCounts; i++ {
				size = size + echan(dataPtr, subElemTypePtr)
			}
		case kindFunc:
		case kindInterface:
			for i := 0; i < sliceCounts; i++ {
				size = size + eface(dataPtr, subElemTypePtr)
			}
		case kindMap:
			for i := 0; i < sliceCounts; i++ {
				size = size + emaps(dataPtr, subElemTypePtr)
			}
		case kindPtr:
			for i := 0; i < sliceCounts; i++ {
				size = size + eptrs(dataPtr, subElemTypePtr)
			}
		case kindSlice:
			subDataPtrs := *(*unsafe.Pointer)(dataPtr)
			for i := 0; i < sliceCounts; i++ {
				size = size + eslice(unsafe.Add(subDataPtrs, uintptr(i)*word*3), subElemTypePtr)
			}
		case kindString:
			subDataPtrs := *(*unsafe.Pointer)(dataPtr) // slice having structure
			l := sliceCounts * 2
			for i := 0; i < l; i = i + 2 {
				size = size + *(*int)(unsafe.Add(subDataPtrs, uintptr(i+1)*word))
			}
		case kindStruct:
			for i := 0; i < sliceCounts; i++ {
				size = size + estruct(dataPtr, subElemTypePtr)
			}
		case kindUnsafePointer:
			for i := 0; i < sliceCounts; i++ {
				size = size + eunptr(dataPtr, subElemTypePtr)
			}
		default:
		}
	}
	return size
}

type slicetype struct {
	typ  _type
	elem *slicetype
}
