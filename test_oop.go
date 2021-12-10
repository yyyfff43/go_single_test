package main

import (
	"fmt"
	"strconv"
)
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

//声明一个叫做Rect的结构体
type Rect struct {
	x, y float64
	width, height float64
}
//添加一个Rect的成员方法
func (r *Rect) Area() float64 {
	return r.width * r.height
}

//声明一个数学算是接口
type IShuXue interface {
	JiaFa(x int, y int) (str_sum string, err error)
}

//声明第一个结构体实现数学接口,实现正整数相加
type ShiXianJiaFa struct {

}

func (s *ShiXianJiaFa) JiaFa(x int, y int) (str_sum string, err error){
	return "这是第一个实现数学接口方法的结构体,正整数相加结果为"+strconv.Itoa(x+y), nil
}

//声明第二个结构体实现数学接口,实现负整数相加
type ShiXianJiaFa_2 struct {

}

func (s *ShiXianJiaFa_2) JiaFa(x int, y int) (str_sum string, err error){
	return "这是第二个实现数学接口方法的结构体,负整数相加结果为"+strconv.Itoa(-x + -y), nil
}

func main (){
	fmt.Println("测试结构体")
	//实例化Rect的对象，一共四种
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 50, 100}
	rect4 := &Rect{width: 100, height: 200}

	fmt.Println("第一种声明调用Area方法为：%f",rect1.Area())
	fmt.Println("第二种声明调用Area方法为：%f",rect2.Area())
	fmt.Println("第三种声明调用Area方法为：%f",rect3.Area())
	fmt.Println("第四种声明调用Area方法为：%f",rect4.Area())

	/*
	在Go语言中，未进行显式初始化的变量都会被初始化为该类型的零值，例如bool类型的零
    值为false，int类型的零值为0，string类型的零值为空字符串。
    在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名，表示“构造函数”：
    func NewRect(x, y, width, height float64) *Rect {
          return &Rect{x, y, width, height}
	}
	 */

	 //接口学习开始:
	 //第一种接口实现，实现正数相加
     var sxjf IShuXue = new(ShiXianJiaFa)
     fmt.Println(sxjf.JiaFa(1,2))
     //第二种接口实现，实现负数相加
	 var sxjf_2 IShuXue = new(ShiXianJiaFa_2)
	 fmt.Println(sxjf_2.JiaFa(1,2))
}

/*
接口(非侵入式)
在Go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口，
例如：
type File struct {
 // ...
}
func (f *File) Read(buf []byte) (n int, err error){方法内部功能省略……}
func (f *File) Write(buf []byte) (n int, err error){方法内部功能省略……}
func (f *File) Seek(off int64, whence int) (pos int64, err error){方法内部功能省略……}
func (f *File) Close() error
这里我们定义了一个File类，并实现有Read()、Write()、Seek()、Close()等方法。设
想我们有如下接口：
type IFile interface {
 Read(buf []byte) (n int, err error)
 Write(buf []byte) (n int, err error)
 Seek(off int64, whence int) (pos int64, err error)
 Close() error
}
type IReader interface {
 Read(buf []byte) (n int, err error)
}
type IWriter interface {
 Write(buf []byte) (n int, err error)
}
type ICloser interface {
 Close() error
}
尽管File类并没有从这些接口继承，甚至可以不知道这些接口的存在，但是File类实现了
这些接口，可以进行赋值：
var file1 IFile = new(File)
var file2 IReader = new(File)
var file3 IWriter = new(File)
var file4 ICloser = new(File)
(如果有另外的File2类等实现了这些同名接口但不同方法内部的功能，均可以赋值给这些对象，实现接口的多钟实例化)
 */
