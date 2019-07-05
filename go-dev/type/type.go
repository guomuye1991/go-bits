package main

import "fmt"

func main() {
	//Demo1()
	Demo2()
}

type Mybit = byte //别名

/*
引用类型
*/
func Demo1() {
	m := make(map[string]int) //new(map[string]int)
	c := make(chan int)
	s := make([]string, 10)
	fmt.Println(m, c, s)
	m["gmy"] = 1
}

/*
类型转换
*/
func Demo2() {
	//类型别名无需显式类型转换
	var b1 Mybit = 1
	var fl byte = b1
	fmt.Println(b1, fl)
	//常量无需类型转换
	const a float64 = 1
	//const b int = 1.1 运行时报错
	var aaa = 1 + 6i
	fmt.Println("%T", aaa)
}

/*
指针，单向channel，无返回值函数类型转换语法歧义问题
*/
func Demo3() {
	i := 0
	p := (*int)(&i)
	fmt.Println(p)
}
