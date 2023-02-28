package std_log

func LogTest() {
	test1()
}

func test1() {
	// print 单纯打印日志
	// panic 打印日志 ， 抛出panic异常，会执行defer
	// fatal 打印日志 ，强制结束程序（os.Exit(1)）, defer 函数不会执行

}
