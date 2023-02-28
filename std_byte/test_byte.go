package std_byte

import (
	"bytes"
	"fmt"
	"strings"
)

func TestByte() {

	test1()
	testContain()
	testCount()
	testRepeat()
	testReplace()
	testRunes()
	testJoin()
}

func test1() {
	// 字符串和字节切片互转
	s := "xxxkkkk"
	b := []byte(s)
	fmt.Printf("b: %v\n", string(b))

	b1 := []byte{'x', 'f'}
	s1 := string(b1)
	fmt.Printf("s1: %v\n", s1)
}

func testContain() {
	// 字符串和字节切片的包含关系
	b := strings.Contains("hello world", "hello")
	b2 := strings.Contains("xxxx", "yyy")
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b2: %v\n", b2)

	byte1 := []byte("xxdfsadf")
	byte2 := []byte("xxd")
	b3 := bytes.Contains(byte1, byte2)
	fmt.Printf("b3: %v\n", b3)

}

func testCount() {
	// 统计字符串出现次数
	s := []byte("helloooooooooooooo")
	b := []byte("e")
	b2 := []byte("l")
	b3 := []byte("o")

	fmt.Printf("bytes.Count(s, b): %v\n", bytes.Count(s, b))
	fmt.Printf("bytes.Count(s, b2): %v\n", bytes.Count(s, b2))
	fmt.Printf("bytes.Count(s, b3): %v\n", bytes.Count(s, b3))
}

func testRepeat() {
	s := []byte("abc")
	fmt.Printf("bytes.Repeat(s, 4): %v\n", string(bytes.Repeat(s, 4)))
}

func testReplace() {
	s := []byte("hello world")
	old := []byte("o")
	new := []byte("ee")
	// replace the third param means the count of replacement , if it's <0 means all
	fmt.Printf("bytes.Replace(old, new, 1): %v\n", string(bytes.Replace(s, old, new, 1)))
	fmt.Printf("bytes.Replace(old, new, 2): %v\n", string(bytes.Replace(s, old, new, 2)))
	fmt.Printf("bytes.Replace(old, new, -1): %v\n", string(bytes.Replace(s, old, new, -1)))
}

func testRunes() {
	s := []byte("你好世界！")
	r := bytes.Runes(s)
	// 一个汉子三个字节长度，需要先转化为Runes来计算长度
	fmt.Printf("string(s): %v\n", string(s))
	fmt.Printf("len(string(s)): %v\n", len(string(s)))
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("len(r): %v\n", len(r))
}

func testJoin() {
	s := [][]byte{[]byte("你好"), []byte("世界")}
	setp := []byte(",")
	fmt.Printf("string(bytes.Join(s, setp)): %v\n", string(bytes.Join(s, setp)))
	setp1 := []byte("$")
	fmt.Printf("string(bytes.Join(s, setp)): %v\n", string(bytes.Join(s, setp1)))
}
