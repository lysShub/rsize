package rsize

import (
	"fmt"
	"unsafe"
)

const structfieldOffset uintptr = 3 * word

func estruct(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {
	fieldsDataPtr := (unsafe.Add(typePtr, typeOffsed+word))
	// *(*unsafe.Pointer)
	a := (*[]Structfield)(fieldsDataPtr)
	fmt.Println(a)

	fieldsLen := *(*uintptr)(unsafe.Add(typePtr, typeOffsed+word+word)) // slice's len field

	var dataOffset uintptr
	for i := uintptr(0); i < fieldsLen; i++ {
		fieldTypePtr := *(*unsafe.Pointer)(unsafe.Add(fieldsDataPtr, i*structfieldOffset+word))
		offsetAnon := *(*uintptr)(unsafe.Add(fieldsDataPtr, i*structfieldOffset+word*2))

		b := (*_type)(fieldTypePtr)
		fmt.Println(b, offsetAnon)

		fieldKind := *(*uint8)(unsafe.Add(fieldTypePtr, 2*word+7))
		dataOffset = dataOffset + *(*uintptr)(unsafe.Add(fieldTypePtr, 0))

		if tsize := originKind(fieldKind); tsize != 0 {
			size = size + tsize
		} else {
			fieldDataPtr := unsafe.Add(dataPtr, dataOffset)
			switch fieldKind {
			case kindArray:
				size = size + earray(fieldDataPtr, fieldTypePtr)
			case kindChan:
			case kindFunc:
			case kindInterface:
			case kindMap:
				size = size + emaps(fieldDataPtr, fieldTypePtr)
			case kindPtr:
			case kindSlice:
				size = size + eslice(fieldDataPtr, fieldTypePtr)
			case kindString:
				size = size + *(*int)(unsafe.Add(fieldDataPtr, word))
			case kindStruct:
				size = size + estruct(fieldDataPtr, fieldTypePtr)
			case kindUnsafePointer:
			default:
			}

		}
	}
	return
}

type Structtype struct {
	typ     _type
	pkgPath name
	fields  []Structfield
}

type Structfield struct {
	name       name
	typ        *_type
	offsetAnon uintptr
}

type name struct {
	bytes *byte
}
