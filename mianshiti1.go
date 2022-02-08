/*
* @File : mianshiti1
* @Describe : 面试题练习1
* @Author: Jerry Yang
* @Date : 2022/2/8 23:58
* @Software: GoLand
 */

package main

import "fmt"

//递归：函数自己调用自己
//递归函数一定要一个明确的退出
//递归适合处理那种问题相同且规模越来越小的场景

//递归函数实现阶乘算法
//3! = 3*2*1 = 3*2!
//4! = 4*3*2*1 = 4*3!
//5! = 5*4*3*2*1 = 5*4!
func f(n uint64) uint64{
	if n<=1{
		return 1
	}
	return n * f(n-1)
}

//上台阶面试题（也是递归实现）
//上台阶，可以一次上一步或者一次上两部，按照指定的台阶数n算出有多少种方法
func  shangTaiJie(n uint64) uint64  {
	if n==1{//当只有一级台阶时一次只能上一步，一种走法
		return 1
	}
	if n==2{//当有两级台阶时一次上一步走两次或者一次上两步直接走完，两种走法
		return 2
	}
	return shangTaiJie(n-1) + shangTaiJie(n-2)
}

func main() {
	//调用算阶乘
	ret := f(8)
	fmt.Println(ret)

	fmt.Println("")

	//调用上台阶算法
	retShangTaiJie := shangTaiJie(3)//三阶都是每次迈一部步+第一阶迈一步，第二三阶迈一步+第一二阶迈一步，第三阶迈一步
	fmt.Println(retShangTaiJie) //结果为3

}