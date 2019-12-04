package main

import "fmt"

func main() {
	//Demo1()
	//Demo2()
	Demo4()
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

/*
类型一致规则
*/
func Demo4() {
	//指针
	x, y := 1, 1.1
	px, py := &x, &y
	//px = py 编译器提示类型问题
	fmt.Println(px, py)

	//数组
	xarr, yarr, zarr := [8]int{}, [8]float64{}, [2]int{}
	//xarr = yarr
	//xarr = zarr
	fmt.Println(xarr, yarr, zarr)

	//切片
	xslice, yslice, zslice := make([]int, 1), make([]float64, 1), make([]int, 2)
	//xslice = yslice
	xslice = zslice
	fmt.Println(xslice, yslice, zslice)

	//字典
	xmap, ymap, zmap, mmap := make(map[string]int), make(map[float64]int), make(map[string]int), make(map[string]string)
	//xmap = ymap
	xmap = zmap
	//xmap = mmap
	fmt.Println(xmap, ymap, zmap, mmap)

	//通道
	xch, ych, zch, mch, nch := make(chan int), make(chan string), make(chan int, 1), make(<-chan int), make(chan<- int)
	//xch = ych
	xch = zch
	//xch = mch
	//xch = nch
	mch, nch = xch, xch //双向管道可以赋值给任何单向管道
	fmt.Println(xch, ych, zch, mch, nch)

	//结构体
	xst, yst, zst, mst, nst, ost := struct {
	}{}, struct {
	}{}, struct {
		name string
		age  int
	}{}, struct {
		age  int
		name string
	}{}, struct {
		name string
		age  float64
	}{}, struct {
		name string `json:"showName"`
		age  int
	}{}
	xst = yst
	//xst = zst
	//zst = mst
	//zst = nst
	//zst = ost
	fmt.Println(xst, yst, zst, mst, nst, ost)

	//函数
	xf, yf, zf, mf, nf := func() {}, func() {}, func(a int) {}, func() int { return 0 }, func(a int) float64 { return 0 }
	xf = yf
	//xf = zf
	//xf = mf
	//xf = nf
	xf = test1 //函数名无关!!!
	fmt.Println(xf, yf, zf, mf, nf)

	//接口,自己验证吧

}

func test1() {

}
