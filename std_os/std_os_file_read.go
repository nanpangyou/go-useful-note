package std_os

import (
	"fmt"
	"io"
	"os"
)

func StdOsFileStruct() {
	fmt.Println("==================================================")
	fmt.Println("std os lib file struct")
	fmt.Println("==================================================")
	// openfile()
	// createfile()
	// readops()
	// read_at_function()
	// read_dir_function()
	read_seek_func()
}

func openfile() {
	// 以只读形式打开文件
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		fmt.Printf("f: %v\n", f)
		f.Close()
	}
	// 打开还可以用OpenFile ,参数 os.O_CREATE 如果文件不存在就创建
	f2, err2 := os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("err2: %v\n", err2)
		panic(err2)
	} else {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
		fmt.Printf("f2: %v\n", f2)
		f2.Close()
	}
}

func createfile() {
	// 创建文件
	// 等价于 os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}

	// 在系统的临时目录里创建临时文件
	// 第一个参数是 目录默认： TEMP 第二个参数为文件的文件名前缀
	f2, err2 := os.CreateTemp("", "temp")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("f2: %v\n", f2)
		fmt.Printf("f2.Name(): %v\n", f2.Name())
	}
}

func readops() {
	// 读文件
	// f, _ := os.Open("a.txt")
	// defer f.Close()
	// buf := make([]byte, 3)
	// n, err := f.Read(buf)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// } else {
	// 	fmt.Printf("n: %v\n", n)
	// 	fmt.Printf("buf: %v\n", string(buf))
	// }
	// 读取全部文件
	f, _ := os.Open("a.txt")
	defer f.Close()
	for {
		// buffer 定义在循环体内，不然的话最后一次可能会有上次循环的残留
		buf := make([]byte, 10)
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("buf: %v\n", string(buf))
	}
}

func read_at_function() {
	f, _ := os.Open("a.txt")
	defer f.Close()
	buf := make([]byte, 5)
	// 从偏移量开始读
	n, err := f.ReadAt(buf, 3)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
	}
}

func read_dir_function() {
	// 读取(遍历)目录
	f, err := os.Open("test1")
	defer f.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {

		//Note that Readdir takes an argument n, which specifies the maximum number of entries to return. A negative value means to return all entries.
		de, _ := f.ReadDir(-1)
		for _, v := range de {
			fmt.Printf("v.IsDir(): %v\n", v.IsDir())
			fmt.Printf("v.Name(): %v\n", v.Name())
		}
		fmt.Printf("f: %v\n", f)
	}

}

func read_seek_func() {
	f, _ := os.Open("a.txt")
	defer f.Close()
	// 设置偏移量 seek返回当前偏移量和错误
	// seek的第一个参数为偏移量，第二个参数为偏移量的起始位置
	// 第二个参数值：  0 为文件开头 1 为当前位置 2 为文件结尾
	ret, _ := f.Seek(3, 1)
	fmt.Printf("ret: %v\n", ret)
	//再次设置seek（ret,1） 相当于readat(buf , 6)
	f.Seek(ret, 1)
	buf := make([]byte, 5)
	n, err := f.Read(buf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("n: %v\n", n)
		fmt.Printf("buf: %v\n", string(buf))
	}

}
