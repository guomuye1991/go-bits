# Var
1. 静态类型语言
1. 可以改变变量值，但无法改变变量地址
1. 编译后变量仅仅是个内存地址，内存地址存储一个`值`
1. Go中任意变量未显式初始化都会被赋值`零值`

## 变量
> `var`关键字，支持多种语法糖，并且`compiler`会根据上下文推断类型
1. `var v1,v2,...vn {type} {= 实参1,实参2,...实参n}`：`type`和`实参`二者至少存在一方
    * `type`和`实参`同时存在=>类型和值所属必须一致
    * 只`type`=>声明了一个变量，并且编译器会默认赋值`零值`
    * 只 `实参`=>声明了一个与实参类型一致的变量，并且值是`实参`;多变量声明时，实参部位的多个实参值类型可以不一样
1. `var (
        v1...组合语法糖，所有变量声明提取`var`关键字到外围，并且`()`环绕变量体，这种语法糖可以实现`分组`的概念，这种写法不影响作用域
    )` 
1. `v1,v2,...vn:= 实参1,实参2,...实参n`：短语形式只能出现在`func`中，除了声明变量外还能退化为赋值(参考`go-dev/variable/var.go#Demo1`)
1. 变量赋值总是先计算完右侧所有值后才会执行赋值，经典案例(参考`go-dev/variable/var.go#Demo2`)
1. 未使用的局部遍历，编译器是不通过的

## 常量
>常量`const`将上面的`var`替换下就是常量声明;运行时不可变
1. 可以是一个值
1. 也可以是编译时可计算的表达式(常量表达式)
1. `const(
        x = 1
        y
    )`写法，声明y可以无赋值，无类型，这样代表当前变量与前一个变量赋一样的值，这种写法不影响作用域(参考`go-dev/variable/var.go#Demo3`)
1. 局部常量可以只声明，不使用

## 枚举
>Go没有枚举，但是可以使用`iota`常量计数器实现枚举类型，(参考`go-dev/variable/var.go#Demo4`)
1. 自增返回值默认类型是`int`，但是本质上返回时是一个无类型整数
1. iota作用于`const c = iota`：当前计数器仅对c有一次计数c=0，如果再声明`const d = iota`，这代表另一个常量计数器d=0；因此作用于单个常量时只对单个常量生效
1. iota作用于`const()`时，对组内所有常量生效，并且每编译常量组内一个成员,iota就会自增
1. 一个常量组内的声明多路`iota`，但是底层共享的是一个计数器

### 常量变量区别
1. 运行时不变
1. 不存在指针，编译器会把常量在编译阶段所有使用常量的位置全部展开为字面量=>称为`常量展开`(参考`go-dev/variable/var.go#Demo5`)