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
	//ch <- value 写入数据， 向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据。
	//value := <-ch 读出数据，如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel中被写入数据为止。

	/*
	 select 异步管理channel IO操作

		select {
			case <-chan1:
			// 如果chan1成功读到数据，则进行该case处理语句
			case chan2 <- 1:
			// 如果成功向chan2写入数据整数1，则进行该case处理语句
			default:
			// 如果上面都没有成功，则进入default处理流程
		}
	 */

	 /*
	  带缓冲机制的channel
	 c := make(chan int, 1024)
在调用make()时将缓冲区大小作为第二个参数传入即可，比如上面这个例子就创建了一个大小
为1024的int类型channel，即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被
填完之前都不会阻塞。可以使用range来遍历
	 for i := range c {
        fmt.Println("Received:", i)
	 }

	 超时机制：在并发编程的通信过程中，最需要处理的就是超时问题，即向channel写数据时发现channel
已满，或者从channel试图读取数据时发现channel为空。如果不正确处理这些情况，很可能会导
致整个goroutine锁死。
     Go语言没有提供直接的超时处理机制，但我们可以利用select机制。虽然select机制不是
专为超时而设计的，却能很方便地解决超时问题。因为select的特点是只要其中一个case已经
完成，程序就会继续往下执行，而不会考虑其他case的情况。
基于此特性，我们来为channel实现超时机制：
        // 首先，我们实现并执行一个匿名的超时等待函数
        timeout := make(chan bool, 1)
		go func() {
		 time.Sleep(1e9) // 等待1秒钟
		 timeout <- true
		}()
		// 然后我们把timeout这个channel利用起来
		select {
		 case <-ch:
		 // 从ch中读取到数据
		 case <-timeout:
		 // 一直没有从ch中读取到数据，但从timeout中读取到了数据，即超时1秒后执行这个timeout case


	 在Go语言中channel本身也是一个原生类型，与map之类的类型地位一样，因
	此channel本身在定义后也可以通过channel来传递。
    例：
	 type PipeData struct {
		 value int
		 handler func(int) int
		 next chan int
     }
    然后我们写一个常规的处理函数。我们只要定义一系列PipeData的数据结构并一起传递给
    这个函数，就可以达到流式处理数据的目的：
	func handle(queue chan *PipeData) {
	 for data := range queue {
	 data.next <- data.handler(data.value)
	 }
	}

    单向channel：所谓的单向channel概念，其实只是对channel的一种使用限制。
	 var ch1 chan int // ch1是一个正常的channel，不是单向的
     var ch2 chan<- float64// ch2是单向channel，只用于写float64数据
     var ch3 <-chan int // ch3是单向channel，只用于读取int数据

	 channel是一个原生类型，因此不仅支持被传递，还支持类型转换，
	 类型转换对于channel的意义：就是在单向channel和双向channel之间进行转换
	 ch4 := make(chan int)  //ch4 是一个普通的channel
     ch5 := <-chan int(ch4) //约定限制ch4后，ch5就是一个单向的读取channel
     ch6 := chan<- int(ch4) //约定限制ch4后，ch6 是一个单向的写入channel
     这样设计的约定遵循“最小权限原则”

	 关闭channel ：调用close()函数即可
	 close(ch)

	 	在介绍了如何关闭channel之后，我们就多了一个问题：如何判断一个channel是否已经被关
	 闭？我们可以在读取的时候使用多重返回值的方式：
	 x, ok := <-ch
	 这个用法与map中的按键获取value的过程比较类似，只需要看第二个bool返回值即可，如
	 果返回值是false则表示ch已经被关闭。

	 多核调用：
	 goroutine是通过直接设置环境变量GOMAXPROCS的值，或者在代码中启动goroutine之前先调用以下这个语句以设置使用16个CPU核心：
        runtime.GOMAXPROCS(16)
	 到底应该设置多少个CPU核心呢，其实runtime包中还提供了另外一个函数NumCPU()来获取核心数
	 后续go版本会陆续抛弃GOMAXPROCS然后自行决定如何分配多核

	 可以在每个goroutine中控制何时主动出让时间片给其他goroutine，这可以使用runtime
包中的Gosched()函数实现
*/


/*
	 锁的类型Mutex和RWMutex，RWMutex相对友好些，是经典的单写多读模型。在读锁占用的情
况下，会阻止写，但不阻止读，也就是多个goroutine可同时获取读锁（调用RLock()方法；而写
锁（调用Lock()方法）会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine
独占。
	 对于这两种锁类型，任何一个Lock()或RLock()均需要保证对应有Unlock()或RUnlock()
调用与之对应，否则可能导致等待该锁的所有goroutine处于饥饿状态，甚至可能导致死锁。锁的
典型使用模式如下：
	var l sync.Mutex
	func foo() {
	 l.Lock()
	 defer l.Unlock()
	 //...
	}
	 调用foo时，如果不能保证在程序末尾执行关闭l.Unlock()，那么最好开头用defer声明
*/

/*
	对于从全局的角度只需要运行一次的代码，比如全局初始化操作，Go语言提供了一个Once
	类型来保证全局的唯一性操作

	var a string
	var once sync.Once
	func setup() {
		a = "hello, world"
	}
	func doprint() {
		once.Do(setup)
		print(a)
	}
	func twoprint() {
		go doprint()
		go doprint()
	}
once的Do()方法可以保证在全局范围内只调用指定的函数一次（这里指setup()函数），而且所有其他goroutine在调用到此语句时，将会先被阻塞，直至全局唯一的
once.Do()调用结束后才继续。

为了更好地控制并行中的原子性操作，sync包中还包含一个atomic子包，它提供了对于一
些基础数据类型的原子操作函数
*/


}
