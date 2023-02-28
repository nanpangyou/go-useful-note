package std_bufio

import (
	"bufio"
	"fmt"
	"os"
)

func Testbufio() {
	test1()
}

func test1() {
	// 任意一个reader
	// r := strings.NewReader("xxxxx")

	f, _ := os.Open("a.txt")
	defer f.Close()
	// bufio 的reader可以封装任意一个 其他reader
	bufReader := bufio.NewReader(f)
	s, _ := bufReader.ReadString('\n')
	fmt.Printf("s: %v\n", s)
}
