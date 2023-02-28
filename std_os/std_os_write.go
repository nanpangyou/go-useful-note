package std_os

import (
	"fmt"
	"os"
)

func OsWrite() {
	// 写相关操作
	// basic_write()
	// write_with_trunc()
	// write_with_append()
	write_at()
	f, _ := os.OpenFile("b.txt", os.O_CREATE|os.O_RDWR, 0755)
	defer f.Close()
	f.Write([]byte("test"))
}

func basic_write() {
	// 这里不可以使用 os.open() 因为这个是只读的
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0755)
	defer f.Close()
	// 如果只有os.O_RDWR 则会在最开头开始复制你写入的字符串 以下操作会： abcdefg => xxxdefg
	f.Write([]byte("xxx"))
}

func write_with_trunc() {
	// os.O_RDWR|os.O_TRUNC 会覆盖文本
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	defer f.Close()
	f.Write([]byte("replace text"))
}

func write_with_append() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755)
	defer f.Close()
	f.Write([]byte("append text"))
}

func write_at() {
	f, err := os.OpenFile("a.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("err21: %v\n", err)
	}
	defer f.Close()
	// 会从偏移量3开始覆盖写入
	n, err2 := f.WriteAt([]byte("write at"), 3)
	if err2 != nil {
		fmt.Printf("err2233: %v\n", err2)
	} else {
		fmt.Printf("n: %v\n", n)
	}
}
