package main

import "fmt"
/*结构体：
func (a *Integer) Add(b Integer) {
	*a += b
}
这里为Integer类型增加了Add()方法。由于Add()方法需要修改对象的值，所以需要用指针引用。调用如下：
func main() {
	var a Integer = 1
	a.Add(2)
	fmt.Println("a =", a)
}
运行该程序，得到的结果是：a=3。如果你实现成员方法时传入的不是指针而是值（即传入 Integer，而非*Integer），如下所示：
func (a Integer) Add(b Integer) {
	a += b
}
那么运行程序得到的结果是a=1，也就是维持原来的值。读者可以亲自动手尝试一下。
*/

/*
Go语言中的数组和基本类型没有区别，是很纯粹的值类型，例如：

var a = [3]int{1, 2, 3}
var b = a
b[1]++
fmt.Println(a, b)
该程序的运行结果如下：
[1 2 3] [1 3 3]。

这表明b=a赋值语句是数组内容的完整复制。要想表达引用，需要用指针：

var a = [3]int{1, 2, 3}
var b = &a
b[1]++
fmt.Println(a, *b)
该程序的运行结果如下：
[1 3 3] [1 3 3]
这表明b=&a赋值语句是数组内容的引用。变量b的类型不是[3]int，而是*[3]int类型
*/
func main (){
	fmt.Println("测试结构体")
}
