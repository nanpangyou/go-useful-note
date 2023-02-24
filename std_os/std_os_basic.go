package std_os

import (
	"fmt"
	"os"
)

func StdOsLibBasic() {
	fmt.Println("==================================================")
	fmt.Println("std os lib")
	fmt.Println("==================================================")
	ch := make(chan struct{})
	go createFile(ch)
	// createDirAll()
	// remove()
	// wd()
	<-ch
	rename()
	writeFile()
	readFile()
}

func createFile(ch chan struct{}) {
	// 创建文件，但是文件重复的时候会覆盖
	fmt.Printf("创建文件: %v\n", "创建文件")
	f, err := os.Create("a.txt")
	ch <- struct{}{}
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
}

func createDir() {
	// 创建单层文件夹 如果已经存在会报错
	fmt.Printf("创建文件夹: %v\n", "创建文件夹")
	err := os.Mkdir("dir", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

func createDirAll() {
	// 创建多层文件夹
	fmt.Printf("创建文件夹: %v\n", "创建文件夹")
	err := os.MkdirAll("dir/dir3", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

func remove() {
	// 删除目录或者文件
	// Remove removes the named file or (empty) directory
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	// RemoveAll removes path and any children it contains
	err2 := os.RemoveAll("dir")
	if err2 != nil {
		fmt.Printf("err: %v\n", err2)
		panic(err2)
	}
}

func wd() {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}

	// 更换目录  貌似不能用家目录（～）开头
	err2 := os.Chdir("/Users/nanpangyou/Code")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	// 获取当前系统的临时目录
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func rename() {
	err := os.Rename("a.txt", "b.txt")
	if err != nil {
		fmt.Printf("err3: %v\n", err)
	}
}

func writeFile() {
	err := os.WriteFile("b.txt", []byte("hello world111"), os.ModePerm)
	if err != nil {
		fmt.Printf("err2: %v\n", err)
	}
}

func readFile() {
	b, err := os.ReadFile("b.txt")
	if err != nil {
		fmt.Printf("err1: %v\n", err)
	} else {
		fmt.Printf("b: %v\n", string(b[:]))
	}
}
