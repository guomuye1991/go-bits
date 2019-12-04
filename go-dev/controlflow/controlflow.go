package main

import "fmt"

/*
if...else...
*/
func main() {
	//Demo1()
	//Demo3()
	Demo4()
}

func Demo1() {
	if initIf(); true {
		fmt.Println(true)
	}
}

func initIf() {
	fmt.Println("initIf")
}

/*
switch
*/
func Demo2() {
	x := 1
	switch x {
	case 1, 2, 3:
		fmt.Println(1, 2, 3)
	case 4:
		fmt.Println(4)
	//case 1: 冲突
	default:
		fmt.Println("default")
	}

	//相邻空case不构成多条件匹配
	switch x {
	case 1: //这里会被插入代码break
	case 2:

	}
}

func Demo3() {
	x := 1
	switch x {
	case 1:
		fmt.Println("1")
		fallthrough //必须写到最后一行
	case 2: //会被执行虽然x不是2
		fmt.Println("2")
	case 3: // 不会被执行，如果想要执行必须在case2继续追加fallthrough
		fmt.Println("3")
	default:
		fmt.Println("default")

	}
}

func Demo4() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i < 10 {
		i++
		fmt.Println(i)
	}
	sl := make([]int, 2)
	sl[0], sl[1] = 0, 1
	//for...range
	for i, v := range sl {
		fmt.Println(i, v)
	}
}

/*
goto
*/
func Demo5() {
	a, b := 1, 1
	fmt.Println(a, b)
	if 2 == a+b {
		goto B
	} else {
		goto C
	}

B:
	fmt.Println("2")
C:
	fmt.Println("3")
}
