package main

import "fmt"

func main() {

	//字符串遍历
	str := "Hello,世界"
	fmt.Println(string(str[0]))
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}
	fmt.Println("如果遍历字符串每个元素，打出的是字节数")

	//数组切片
	mySlice := make([]int, 5, 10) //创建一个默认元素都是0的，5个数组元素的，容量为10的切片
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))

	mySlice = append(mySlice, 1, 2, 3) //追加内容为1,2,3的切片
	mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...) //...省略号必须有，追加myslice2切片，打散后续切片元素后合并
	fmt.Println("合并后的切片为:", mySlice)

	//数组切片复制，以容量小的切片为主
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置

	//map
	// PersonInfo是一个包含个人详细信息的类型
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}
	var myMap map[string]PersonInfo
	myMap = make(map[string]PersonInfo, 100) //第三个参数100是标识map的容量为100，是可选参数
	myMap["1234"] = PersonInfo{"1", "Jack", "Room 101,..."}
	//map的查找
	value, ok := myMap["1234"]
	if ok { // 如果找到了，只判断第二个返回的参数是不是ok即可
		// 处理找到的value
		fmt.Println(value.ID)
	}

	//函数
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)

}

/*
	函数，小写字母开头的函数只在本包内可见，大写字母开头的函数才
	能被其他包使用
*/
func MyPrintf(args ...interface{}) { //不定参数args, ...表示语法糖，多个这个类型的参数输入
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// func (file *File) Read(b []byte) (n int, err Error)   方法名Read前的(file *File) 表示此方法是File结构体的成员方法
//defer 关键字用来做资源回收时的调用，比如File类,当在一个函数执行过程中调用panic()函数时，正常的函数执行流程将立即终止，但函数中
//之前使用defer关键字延迟执行的语句将正常展开执行,另外recover()函数执行将使得该错误处理过程终止


