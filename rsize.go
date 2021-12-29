package rsize

import "unsafe"

func GetEfaceSize(e *interface{}) (size int) {
	typ := (*unsafe.Pointer)(unsafe.Pointer(&e))

	// kind in _type
	kind := (*uint8)(unsafe.Pointer(uintptr(*typ) + 23)) // 8 + 8 + 4 + 1 + 1 + 1
	if size = efacegot(*kind); size != 0 {
		return size
	}

	switch *kind {
	case 17: // array

	case 99:
	}

	return 0
}

// 能够直接直接获取到大小的类型
func gotType(kind uint8) (size int) {
	switch kind {
	case 1, 3, 8: // bool int8 uint8
		return 1
	case 4, 9: // int16 uint16
		return 2
	case 5, 10, 13: // int32 uint32 float32
		return 4
	case 6, 11, 14, 15, 19: // int64 uint64 float64 Complex64 func
		return 8
	case 16: // kindComplex128
		return 16
	default:
		return 0
	}
}

func arrayGot()
