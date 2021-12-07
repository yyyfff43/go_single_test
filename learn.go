package main

import (
      "fmt"
)

func main() {
      var (
            num1 int
            num2 int
            num3 int
      )
      num1, num2, num3 = 1, 2, 3
      // 打印函数调用语句。用于打印上述三个变量的值。
      fmt.Println(num1, num2, num3)


      // 变量声明和赋值语句，由关键字var、变量名num、变量类型uint64、特殊标记=，以及值10组成。
      var num uint64 = 65535

      // 短变量声明语句，由变量名size、特殊标记:=，以及值（需要你来填写）组成。
      size := 524280

      // 打印函数调用语句。在这里用于描述一个uint64类型的变量所需占用的比特数。
      // 这里用到了字符串的格式化函数。
      fmt.Printf("类型为 uint64 的整数 %d 需占用的存储空间为 %d 个字节。\n", num, size)

      var testnum int = 123
      var teststr string = "哈哈哈"
      fmt.Printf("mystest int is %d and str is %s\n" , testnum,teststr)


      var numbers2 [5]int
      numbers2[0] = 2
      numbers2[3] = numbers2[0] - 3
      numbers2[1] = numbers2[2] + 5
      numbers2[4] = len(numbers2)
      sum := 11
      // “==”用于两个值的相等性判断
      fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))

      var numbers3 = [5]int{1, 2, 3, 4, 5}
      slice3 := numbers3[2 : len(numbers3)]

      fmt.Println((cap(slice3)))//容量等于底层数组长度减去切片下索引的值的绝对值

      //切片的更多操作
      var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
      slice5 := numbers4[4:6:8]
      length := 2
      capacity := 4
      fmt.Printf("%v, %v\n", length == len(slice5), capacity == cap(slice5))
      slice5 = slice5[:cap(slice5)] //将slice5的长度延展为何其容量相同,slice5数组值实际变为{5,6,7,8}
      slice5 = append(slice5, 11, 12, 13)
      length = 7
      fmt.Printf("%v\n", length == len(slice5))
      slice6 := []int{0, 0, 0}
      copy(slice5, slice6)
      e2 := 0
      e3 := 8
      e4 := 11
      fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])

      //字典类型
      mm2 := map[string]int{"golang": 42, "java": 1, "python": 8}
      mm2["scala"] = 25
      mm2["erlang"] = 50
      mm2["python"] = 0

      e, ok := mm2["hahaha"] //e得到默认值0，ok变量是false，因为键为hahaha的元素不存在，如果存在e得到实际值，ok为true

      fmt.Printf("%d, %d, %d, %v, %v \n", mm2["scala"], mm2["erlang"], mm2["python"],e,ok)

      //通道channel类型
      ch2 := make(chan string, 1)
      // 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
      // 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
      // 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
      // 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
      go func() {
            ch2 <- "已到达！"
      }()
      var value string = "数据"
      value = value + <-ch2
      fmt.Println(value)

//      var myChannel = make(chan int, 0) //非缓冲管道，最后一个参数为0

//      var sender Sender = myChannel //定义单向发送通道
//      var receiver Receiver = myChannel //定义单向接收通道

}