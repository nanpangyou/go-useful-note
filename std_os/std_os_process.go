package std_os

import (
	"fmt"
	"os"
)

func Process() {
	opt_process()

}

func opt_process() {
	// 获取当前进程
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	// 获取当前的父进程
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())

	// os.StartProcess()

	// 其他进程相关的用的的时候看官方文档
}
