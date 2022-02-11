package main

import (
	"fmt"
)

type person struct {
	name string
	city string
	age  int8
}

type student struct {
	name string
	age  int
}

type Person2 struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person2) SetDreams(dreams []string) {
//	p.dreams = dreams //这里是值复制，外部如果调用切片会改变其中的值，因为切片数据类型包含了指向底层数据的指针
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}

	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}

	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	//这个相同于new(person)
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}

	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}

	//面试题，下边输出是什么
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	p222 := Person2{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p222.SetDreams(data)

	data[1] = "不睡觉"//如果知识单纯的把data靠=赋值给p222.dreams，则睡觉会变成不睡觉，如果是用copy方法则实现了值复制，不会改变p222.dreams
	fmt.Println(p222.dreams)  // ?
}
