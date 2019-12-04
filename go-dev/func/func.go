package main

import "fmt"

func main() {
	//Demo1()
	//Demo2()
	//Demo3()
	Demo6()
}

/*
声明了一种函数类型
*/
type MyFunc func()

/*
函数只能与nil比较
*/
func Demo1() {
	fmt.Println(func1 == nil)
	fmt.Println(func1 != nil)
	//fmt.Println(func1 == func2) compile error
}

func Demo2() {
	a := struct {
	}{}

	fmt.Printf("传入前 %p", &a)
	fmt.Println()
	func4(a)
}

func Demo3() {
	y := 1000
	p1 := &y //指定到这时，p与p1动态值一致
	fmt.Printf("p %p p1 %p", &y, p1)
	fmt.Println()
	p2 := &p1
	func5(p2) //执行完毕后p1的动态值变化，并且值页变化，但是与func5的局部变量x指针值一致
	fmt.Printf("p %p p1 %p", &y, p1)

}

func Demo4() {
	func6(1, "1", "2")
}

func Demo5() {
	func7()
}

func Demo6() {
	f := func() {
		x := 1
		x++
		fmt.Println(x)
	}
	f()
	f()
}

func func1() {

}

func func2() {

}

func func3() *int {
	a := 1
	return &a
}

func func4(i interface{}) {
	fmt.Printf("传入后 %p", &i)
}

func func5(p **int) {
	x := 1
	fmt.Printf("x %p", &x)
	fmt.Println()
	*p = &x
}

func func6(a int, str ...string) {

}

func func7() (a int, b string) {
	a = 1
	b = ""
	//return // 非index使用局部变量复制后返回
	return 2, "123" //使用return指定返回值
}
