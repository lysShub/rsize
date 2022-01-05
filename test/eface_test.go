package test

import (
	"testing"

	"github.com/lysShub/rsize"
	"github.com/stretchr/testify/require"
)

func TestEface(t *testing.T) {
	type test struct {
		name  string
		value interface{}
		size  int
	}

	// 各类型组合嵌套, 存在无数种可能, 测试尽量覆盖, 嵌套不会超过3层
	var tests = []test{
		{
			name:  "bool",
			value: [3]bool{true, false, true},
			size:  3,
		},
		{
			name:  "int8",
			value: [3]int8{1, 6, 9},
			size:  3,
		},
		{
			name:  "uint8",
			value: [3]uint8{1, 6, 9},
			size:  3,
		},
		{
			name:  "int16",
			value: [3]int16{6, 44, 339},
			size:  6,
		},
		{
			name:  "array.uint16",
			value: [3]uint16{6, 44, 339},
			size:  6,
		},
		{
			name:  "array.int32",
			value: [3]int32{6, 44, 339},
			size:  12,
		},
		{
			name:  "array.uint32",
			value: [3]uint32{6, 44, 339},
			size:  12,
		},
		{
			name:  "array.float32",
			value: [3]float32{6.6, 4.4, 339.0},
			size:  12,
		},
		{
			name:  "array.int64",
			value: [3]int64{66, 44, 1<<63 - 1},
			size:  24,
		},
		{
			name:  "array.uint64",
			value: [3]uint64{66, 44, 1<<64 - 1},
			size:  24,
		},
		{
			name:  "array.float64",
			value: [3]float64{66, 44, 1<<64 - 1},
			size:  24,
		},
		{
			name:  "array.complex64",
			value: [3]complex64{66 + 5i, 44, 1<<64 - 1},
			size:  24,
		},
		{
			name:  "array.complex128",
			value: [3]complex128{66 + 5i, 44, 1<<64 - 1},
			size:  48,
		},
		{
			name:  "array.array",
			value: [3][2]int{{1, 1}, {0, 1<<63 - 1}, {1, 2}},
			size:  48,
		},
		{
			name:  "array.array.array",
			value: [3][2][2]int{{{1, 1}, {1, 1}}, {{0, 1<<63 - 1}}, {{1, 2}}},
			size:  96,
		},
		{
			name:  "array.slice",
			value: [3][]int{{1, 1}, {0, 1<<63 - 1}, {1, 2}},
			size:  48,
		},
		{
			name:  "array.string",
			value: [3]string{"1", "123", "1234"},
			size:  8,
		},
	}

	for _, test := range tests {
		require.Equal(t, test.size, rsize.Size(&test.value), test.name)
	}

}
