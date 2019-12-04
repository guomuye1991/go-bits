# 控制流

## if...else...
>`if {expression};condition{...}`(参考`go-dev/expression/expression.go#Demo1`)
1. 支持可选的前置expression表达式用于执行一些初始化操作

## switch
>`switch input_expression {
    case x,y...n://case支持多条件,代表or
        ...
    case m:
        ...
    default:
        ...
}`
1. `break`跳出case，结束switch，go的case也无需显示指定break，每个case结束后会自动break
1. 如果想要case执行完毕后执行紧邻的case，则可以在case块结尾插入`fallthrough`关键字，这样就会无条件的执行紧邻`case/default`
    * 想要连续贯通case就连续插入`fallthrough`即可
1. case不可冲突
1. 编译器对switch和if的指令集一样
1. 相邻空case不构成多条件匹配，会隐式插入`break`(参考`go-dev/controlflow/controlflow.go#Demo2`)

## for
>`for {before};conditionn;{after}{...}`：Go没有while，只有for，可以使用`for condition{}`达到while效果

### for...range
>`range 可迭代值/表达式/变量`：range会返回2个参数`index,value:=range`，range每执行一次就会进行一次迭代
1. `for i,v:= range`配合可以完美的迭代元素`for...range`for会一直循环迭代到range中无元素结束
    * 允许单值返回`for i:=range...`：i是索引；`for _,v:=range`：v是值
    * 空迭代`for range...`
1. 最重要的一点`range data`执行时用的是data的复制！！！因此对于数组这种非引用可迭代类型最好使用指针或用切片替代，否则就有双倍性能开销

## goto
>`goto {tag}`：必须定义tag，并且未使用的tag会编译失败，不能跳转到其它函数tag，以及内层{}内的tag

## break
>`break`中断当前for,switch,select

## continue
>`continue`用于for循环
