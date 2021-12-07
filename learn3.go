package main

import (
	"errors"
	"fmt"
)

func deferIt2() {
	for i := 1; i < 5; i++ {
		defer fmt.Print(i)//都会在该函数即将退出那一刻被执行,这就更进一步地保证了资源的及时释放,当一个函数中存在多个defer语句时，它们携带的表达式语句的执行顺序一定是它们的出现顺序的倒序
	}
}

func Addfunc(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { // 假设这个函数只支持两个非负数字的加法
		err = errors.New("Should be non-negative numbers!")//能够善用error接口、errors.New函数和比较操作符==
		return
	}
	return a + b, nil // 支持多重返回值
}

func main() {
	var number int
	if number := 4; 100 > number {//number的标识符被重声明了，作用域只在if语句内
		number += 3
	} else if 100 < number {
		number -= 2
	} else {
		fmt.Println("OK!")
	}
	fmt.Println(number)//由于在if条件时被重声明，number出了if语句后还是原初始值0

	//使用for循环遍历字典，遍历的结果索引可能是无序的
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	for i, v := range map1 {
		fmt.Printf("%d: %s\n", i, v)
	}

	deferIt2()

}
