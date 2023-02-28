package std_builtin

import "fmt"

func TestBuiltin() {

	testMake()
	testAppend()
	testLen()
	testPanic()
}

func testMake() {
	// new 和 make 的区别
	// 1. make只能用来分配及初始化类型为 `slice`, `map`, `chan` 的数据; new可以分配任意类型的数据
	// 2. new 分配返回的是指针， 即类型 `*T`, make 返回的是引用, 即 `T`
	// 3. new 分配的空间被清零，make分配后会进行初始化

	// new
	b := new(string)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %v\n", *b)

	num := new(int)
	fmt.Printf("num: %T\n", num)
	fmt.Printf("num: %v\n", *num)

	a := new(bool)
	fmt.Printf("a: %T\n", a)
	fmt.Printf("a: %v\n", *a)

	// make
	var p *[]int = new([]int)
	fmt.Printf("p: %v\n", p)

	v := make([]int, 10)
	fmt.Printf("v: %v\n", v)
}

func testAppend() {
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("i: %v\n", i)
	s1 := []int{2, 3, 4}
	i2 := append(s, s1...)
	fmt.Printf("i2: %v\n", i2)
}

func testLen() {
	s := "hello xxx"
	fmt.Printf("len(s): %v\n", len(s))
	s1 := []int{213, 4, 5, 4}
	fmt.Printf("len(s1): %v\n", len(s1))
}

func testPanic() {
	defer fmt.Println("panic也会执行defer")
	panic("抛出panic异常，程序退出")
	fmt.Println("在Panic后，不会执行")
}
