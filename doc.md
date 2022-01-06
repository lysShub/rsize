### 写给自己

1.  常规操作

| N                | normal  | unsafe                  | note                                                         |
| ---------------- | ------- | ----------------------- | ------------------------------------------------------------ |
| 取地址, 得到指针 | p := &a | p := unsafe.Pointer(&a) | unsafe.Pointer不是函数，是强制类型转换                       |
| 取指针的值       | *p      | *(\*unsafe.Pointer)(p)  | 完成两部操作：强制类型转换为unsafe.Pointer，取unsafe指针的值 |
|                  |         |                         |                                                              |



2. 

```go
type a struct {
	v1 uint8
	v2 uint8
}

type A struct {
	a
	p1 a
	p2 *a
}

func main() {
	var v = A{}
	v.v1 = 1
	v.v2 = 2
	v.p1 = a{6, 7}
	v.p2 = &a{11, 13}

	ptr := unsafe.Pointer(&v)

	k1 := *(*uint8)(unsafe.Add(ptr, 0)) // v1 1
	k2 := *(*uint8)(unsafe.Add(ptr, 1)) // v2 2
	k3 := *(*uint8)(unsafe.Add(ptr, 2)) // p1.v1 6
	k4 := *(*uint8)(unsafe.Add(ptr, 3)) // p1.v2 7

	p2ptr := *(*unsafe.Pointer)(unsafe.Add(ptr, 8)) // memory alignment

	k5 := *(*uint8)(unsafe.Add(p2ptr, 0)) // p2.v1 11
	k6 := *(*uint8)(unsafe.Add(p2ptr, 1)) // p2.v23 13

	fmt.Println(k1, k2, k3, k4, k5, k6)
}
```

