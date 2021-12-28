package main

import (
	"encoding/json"
	"fmt"
)

//json数据交换格式学习
type Book struct {
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float64
}

//定义一个空接口，给未知结构json解码
var IrJson interface{}

func main() {
	fmt.Println("使用json.Marshal()函数,json编码:")

    gobook := new(Book)
	gobook = &Book{"Go语言编程",
		[]string{"zhangsan", "lisi", "wangwu", "zhaoliu", "liubei", "guanyu", "zhangfei"},
		"ituring.com.cn",
		true,
		9.99,
	}
	b, err := json.Marshal(gobook)
	//如果编码成功，err 将赋于零值 nil，变量b 将会是一个进行JSON格式化之后的[]byte类型
    if err == nil{
    	fmt.Println(b)
	}else{
		fmt.Println("编码未成功")
	}

	fmt.Println("使用json.Unmarshal()函数，json解码:")
	var outBook Book
	err2 := json.Unmarshal(b, &outBook)
	if err2==nil {
		fmt.Println(outBook)
	}else{
		fmt.Println("解码未成功")
	}

	//json.Unmarshal()函数会根据一个约定的顺序查找目标结构中的字段，如果找到一个即发生匹配。当JSON数据里边的结构和Go里边的目标类型的结构对不上时，
	//如果JSON中的字段在Go目标类型中不存在，json.Unmarshal()函数在解码过程中会丢弃该字段。

	fmt.Println("解码未知结构的JSON数据")
    //要解码一段未知结构的JSON，只需将这段JSON数据解码输出到一个空接口即可

	err3 := json.Unmarshal(b, &IrJson)
	//json.Unmarshal() 函数将一个JSON对象解码到 空接口IrJson中，最终r将会是一个键值对的 map[string]interface{} 结构
	if err3 ==nil {
		fmt.Println(IrJson)
	}

	//要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型：
	unknowBook, ok := IrJson.(map[string]interface{})
	if ok {
		for k, v := range unknowBook {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}
}
