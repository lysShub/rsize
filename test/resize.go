package test

// 测试
// 无法遍历所有类型组合, 因此最多测试至三层嵌套

type t struct {
	Name  string
	Value interface{}
	Size  int
}

var Tests []t = make([]t, 0)

func init() {
	Tests = append(Tests, eorigin...)
	Tests = append(Tests, earray...)
}
