package main

import "fmt"

func main() {
	//Demo1()
	//Demo2()
	//Demo3()
	Demo4()
}

/*
短语声明退化为赋值
*/
func Demo1() {
	x := 100
	fmt.Println("声明一个x变量", &x, x)
	x, y := 300, "abc"
	fmt.Println("赋值给x变量", &x, x)
	fmt.Println("声明一个y变量", &y, y)
	y, z := "cbd", 2.0
	fmt.Println("赋值给y变量", &y, y)
	fmt.Println(z)
}

/*
所有赋值运算都是：先计算完毕所有右侧值，再赋值到左侧
*/
func Demo2() {
	x, y := 1, 2
	fmt.Println("计算前", x, y)
	x, y = y+1, x+3
	fmt.Println("计算后", x, y)
}

func Demo3() {
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)
}

/*
iota定义一个常量计数器，每遇到一个常量都会+1
*/
func Demo4() {
	const (
		a = iota //声明一个iota常量计数器,作用于当前组
		b
		c
		d
		e = 100          //还是会自增
		f                //还是会自增
		g         = iota //使用的是当前组的
		h float64 = iota // 自增返回值默认类型是`int`，但是本质上返回时是一个无类型整数
	)
	fmt.Println(a, b, c, d, e, f, g, h)

	const (
		//a = iota 不影响作用域
		v1, m1 = iota, 100
		v2, m2 = iota, iota * 10 //第二路：iota=1而不是0
		v3, m3
	)
	fmt.Println(v1, m1)
	fmt.Println(v2, m2)
	fmt.Println(v3, m3)
}

func Demo5() {
	const n = 1
	const s = "abc"
	//fmt.Println(&n,&s) 编译失败
}
