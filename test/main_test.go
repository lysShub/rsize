package main_test

import (
	"reflect"
	"testing"
	"unsafe"
)

var is interface{} = "dsafasetfgasfdvasdfasfsad"

func BenchmarkSelf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		typePtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&is)) + 0))
		iKind := (*uint8)(unsafe.Pointer(uintptr(*typePtr) + 23))
		if *iKind != 24 {
			b.Fatal("self kind err")
		}
		strPtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&is)) + 8))

		strLen := (*int)(unsafe.Pointer(uintptr(*strPtr) + 8))
		if *strLen != 25 {
			b.Fatal("self len err")
		}
	}
}

func BenchmarkReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch v := reflect.ValueOf(is); v.Kind() {
		case reflect.String:
			if len(v.String()) != 25 {
				b.Fatal("reflect len err")
			}
		default:
			b.Fatal("reflect kind err")
		}
	}
}
