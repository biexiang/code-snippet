package main

type HugeExplicitT struct {
	a [3 * 1000 * 1000]int32 // 12MB
}

var vCapSize = 10

const cCapSize = 10

func f0() {
	// moved to heap: t1
	t1 := HugeExplicitT{}

	// make([]int32, 0, 17 * 1000) escapes to heap
	t2 := make([]int32, 0, 17*1000)

	// make([]int32, vCapSize, vCapSize) escapes to heap
	t3 := make([]int32, vCapSize, vCapSize)

	// make([]int32, cCapSize, cCapSize) does not escape
	t4 := make([]int32, cCapSize, cCapSize)

	// make([]int32, 10, 10) does not escape
	t5 := make([]int32, 10, 10)

	// make([]int32, vCapSize) escapes to heap
	t6 := make([]int32, vCapSize)

	//  make(map[*int]*int) does not escape
	t7 := make(map[*int]*int)
	// moved to heap: k1
	// moved to heap: v1
	k1, v1 := 0, 0
	t7[&k1] = &v1

	_, _, _, _, _, _ = t1, t2, t3, t4, t5, t6
}

func f1() map[string]int {
	//  make(map[string]int) escapes to heap
	v := make(map[string]int)
	return v
}

func f2() []int {
	// make([]int, 0, 10) escapes to heap
	v := make([]int, 0, 10)
	return v
}

func f3() *int {
	// moved to heap: t
	t := 0
	return &t
}

func f4(v1 *int) **int {
	// moved to heap: v1
	return &v1
}

func main() {
	// 大小超出、切片、映射
	f0()

	// 返回值情况
	_, _, _, _ = f1(), f2(), f3(), f4(&vCapSize)

	var x1 *int
	fn1 := func() {
		// moved to heap: y
		y := 1
		x1 = &y
	}
	fn1()

	var x2 *int
	fn2 := func(input *int) {
		// input does not escape
		y := 1
		input = &y
	}
	fn2(x2)
}
