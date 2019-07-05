# type

## 基本类型
>
### 布尔
1. bool
### 整数
1. byte
1. int
1. int8,int16,int32,int64s
1. uint
1. uint8,uint16,uint32,uint64s
1. rune
1. uintptr
>字面量默认类型是int
### 浮点数
1. float32
1. float64
>字面量默认类型是float64
### 复数
1. complex64
1. complex128
>字面量默认类型是complex128
### 字符串
1. string
### 数组
1. array
### 结构体
1. struct
### 函数
1. function
### 接口
1. interface
### 字典
1. map
### 切片
1. slice
### 同道
1. channel

## 类型别名
> `type {alias_type} = {type}`：alias_type与type拥有同样的底层类型，并且可直接相互赋值；但是有相同的的底层类型不一定可直接复制比如：int在64bit系统上与int64一致，但是不可直接赋值

## 引用类型
>特指`map`,`slice`,`channel`类型，这些类型必须用`make`创建(无法使用new)，编译器会使用特目标类型专用指令初始化
1. `new`也可以为引用类型分配内存，但是这是一种不完善的初始化(仅仅为map分配内存，但是map内部均无分配)，比如`m:=new(map[string]int) m["gmy"]=1`编译报错
1. 参考`go-dev/type/type.go#Demo1`

## 类型转换
>除常量，别名及未命名类型外，Go强制要求显式类型转换`{type}(x)`(参考`go-dev/type/type.go#Demo2`)
1. 如果转换的是指针，单向channel，无返回值函数则需要用`()`否则会语法歧义，无法编译(参考`go-dev/type/type.go#Demo3`)

## 自定义类型
>`type {name} {type}`：基于type构建自定义新类型
1. 语法支持组`type(...)`
1. `{type}`可以是预置的任意类型
1. 与预置类型拥有相同底层结构，享有相同的操作符;但是这是2个完全不一样的类型
    * 不可以直接赋值，比较，隐式转换等
    * 不会继承基础类型的方法等信息
 
 ## 未命名类型
 >有些类型如`bool`,`int`等定义后直接就是确定的类型，但是有些需要与长度，具体类型关联后才能是一个确定类型，如`map,slice,arr,channel,struct,interface,指针`等
 1. 未命名类型无法像确定类型一样直接根据类型名就可判定为同一类型，它们都有各自的规则
 