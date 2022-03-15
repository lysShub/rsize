package test

func init() {
	earray = append(earray, arrayOrigin...)
}

var earray []t

var arrayOrigin = []t{
	{
		Name:  "array.bool",
		Value: [3]bool{true, false, true},
		Size:  3,
	},
	{
		Name:  "array.int8",
		Value: [3]int8{1, 6, 9},
		Size:  3,
	},
	{
		Name:  "array.uint8",
		Value: [3]uint8{1, 6, 9},
		Size:  3,
	},
	{
		Name:  "array.int16",
		Value: [3]int16{6, 44, 339},
		Size:  6,
	},
	{
		Name:  "array.uint16",
		Value: [3]uint16{6, 44, 339},
		Size:  6,
	},
	{
		Name:  "array.int32",
		Value: [3]int32{6, 44, 339},
		Size:  12,
	},
	{
		Name:  "array.uint32",
		Value: [3]uint32{6, 44, 339},
		Size:  12,
	},
	{
		Name:  "array.float32",
		Value: [3]float32{6.6, 4.4, 339.0},
		Size:  12,
	},
	{
		Name:  "array.int64",
		Value: [3]int64{66, 44, 1<<63 - 1},
		Size:  24,
	},
	{
		Name:  "array.uint64",
		Value: [3]uint64{66, 44, 1<<64 - 1},
		Size:  24,
	},
	{
		Name:  "array.float64",
		Value: [3]float64{66, 44, 1<<64 - 1},
		Size:  24,
	},
	{
		Name:  "array.complex64",
		Value: [3]complex64{66 + 5i, 44, 1<<64 - 1},
		Size:  24,
	},
	{
		Name:  "array.complex128",
		Value: [3]complex128{66 + 5i, 44, 1<<64 - 1},
		Size:  48,
	},
	{
		Name:  "array.array",
		Value: [3][2]int{{1, 1}, {0, 1<<63 - 1}, {1, 2}},
		Size:  48,
	},
	{
		Name:  "array.array.array",
		Value: [3][2][2]int{{{1, 1}, {1, 1}}, {{0, 1<<63 - 1}}, {{1, 2}}},
		Size:  96,
	},
	{
		Name:  "array.slice",
		Value: [3][]int{{1, 1}, {0, 1<<63 - 1}, {1, 2}},
		Size:  48,
	},
	{
		Name:  "array.string",
		Value: [3]string{"1", "123", "1234"},
		Size:  8,
	},
}
