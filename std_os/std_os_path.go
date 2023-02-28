package std_os

import (
	"fmt"
	"os"
)

func Path() {
	// 获取所有的环境变量
	fmt.Printf("os.Environ(): %v\n", os.Environ())
	// 根据key-value获取某一个环境变量的值
	fmt.Printf("os.Getenv(\"SHELL\"): %v\n", os.Getenv("SHELL"))
}
