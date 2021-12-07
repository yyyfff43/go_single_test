package main

import (
	"errors"
	"fmt"
)

//函数
func Addfunc(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { // 假设这个函数只支持两个非负数字的加法
		err= errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil // 支持多重返回值，返回值都需要有赋值
}

//结构体
type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Adress  string
}

func (person *Person) Move(newadress string)(oldadress string){//参数Person前必须带指针*，否则调用方法执行的是副本，成员变量和方法的计算结果不会改变
	var oldAdress string = person.Adress
	person.Adress = newadress
	return oldAdress
}


func main() {
	fmt.Println(Addfunc(1, 1))

	//结构体
	p := Person{"Robert", "Male", 33, "Beijing"}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Adress)
}
