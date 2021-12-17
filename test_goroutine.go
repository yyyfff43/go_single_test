package main

import (
	"fmt"
)

//goroutine学习，可配合性能分析包PProf练习
func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	/*
			在一个函数调用前加上go关键字，这次调用就会在一个新的goroutine中并发执行。当被调用
		的函数返回时，这个goroutine也自动结束了。需要注意的是，如果这个函数有返回值，那么这个
		返回值会被丢弃。
	*/
	go Add(1, 1)

	//假设:
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
	/*Go程序从初始化main package并执行main()函数开始，当main()函数返回时，程序退出，
		且程序并不等待其他goroutine（非主goroutine）结束。
		对于上面的例子，主函数启动了10个goroutine，然后返回，这时程序就退出了，而被启动的
		执行Add(i, i)的goroutine没有来得及执行，所以程序没有任何输出。
	    解决办法：
		Go语言提供的消息通信机制被称为channel,不同进程间靠消息来通信，它们不会共享内存。应该通过通信来共享内存。
		channel是Go语言在语言级别提供的goroutine间的通信方式。我们可以使用channel在两个或多个goroutine之间传递消息。
		一个channel只能传递一种类型的值，这个类型需要在声明channel时指定。
	*/
	//定义一个内部传递int的channel
	//	var ch chan int 或
	//	ch := make(chan int) 推荐make
}
