package main

import "fmt"
import "github.com/nonepro/golearn/lang/tree"

// 面向对象只支持封装，不支持继承和多态

// 定义方法
func printValue1(node tree.Node) {
    fmt.Println(node.Value)
}

func main() {

    fmt.Println("内置构建方式")
    var root tree.Node
    root = tree.Node{Value: 3, Left: &tree.Node{}, Right: &tree.Node{}}
    root.Left.Left = &tree.Node{Value: 4}
    root.Left.Right = new(tree.Node)
    root.Right.Value = 5
    root.Right.Left = &tree.Node{Value: 6, Left: &tree.Node{Value: 7}, Right: &tree.Node{Value: 9}}
    fmt.Println("             ", root)
    /******************************************************
     *                      3
     *                0            5
     *            4     0      6
     *                      7     9
     *****************************************************/

    fmt.Println()
    fmt.Println("自定义工厂方法")
    root1 := tree.CreateTreeNode(3)
    fmt.Println("             ", root1)

    fmt.Println()
    fmt.Println("方法声明学习:", "方法的本质就是普通方法,定义方法名和签名")
    fmt.Println("              直接调用printValue1(root)")
    printValue1(root)
    fmt.Println("              问题是局部变量的地址怎么能返回呢？局部变量难道不是在栈空间? => 编译器动态决定堆还是栈")
    fmt.Println("              结论：计算机的基本原理，语言层面不可能改变，语言只能做一些小技巧来运用这些基本原理")

    fmt.Println()
    fmt.Println("包装一下：如何让一组方法相关 => 引入方法接收者概念 (这是工程管理需求，而不是技术实现需求)")
    root.PrintValue2()

    fmt.Println("再进一步：go 到底是传值还是传址")
    fmt.Println("         原始值:", root.Value)
    root.SetValue1(100)
    fmt.Println("         修改后:", root.Value)
    fmt.Println("         结论：结构体也是传值的")

    fmt.Println()
    fmt.Println("再再进一步：值拷贝过于繁琐和可能的性能问题怎么办 => 支持指针类型来传地址的值，这里依然是传值")
    fmt.Println("         原始值:", root.Value)
    root.SetValue(100)
    fmt.Println("         修改后:", root.Value)
    fmt.Println("         传值传址不影响内部代码，都是 node.value = v ?")
    fmt.Println("         >>>有编译器的存在,这些负担不应该留给开发者。虽然解放了手，但是大脑不要偷懒，底层什么逻辑还是要思考")

    fmt.Println()
    fmt.Println("再再再在进一步：")
    fmt.Println("      1、如果用指针类型访问传值结构体的接收者声明呢?")
    root.Value = 0
    pRoot := &root
    fmt.Println("         原始值:pRoot.value =", pRoot.Value)
    fmt.Println("         原始值: root.value =", root.Value)
    pRoot.SetValue1(100)
    fmt.Println("         修改后:pRoot.value =", pRoot.Value)
    fmt.Println("         修改后: root.value =", root.Value)

    fmt.Println("      2、如果用指针类型访问指针类型的接收者声明呢?")
    fmt.Println("         原始值:pRoot.value =", pRoot.Value)
    fmt.Println("         原始值: root.value =", root.Value)
    pRoot.SetValue(100)
    fmt.Println("         修改后:pRoot.value =", pRoot.Value)
    fmt.Println("         修改后: root.value =", root.Value)

    fmt.Println()
    fmt.Println("通过上面的实验，我们发现：")
    fmt.Println("     1、任何语言，方法的本质只是有序指令序列组成的调用过程")
    fmt.Println("     2、区别于其他OO语言this这种重度包装，go只是在形式上引入了函数接收者")
    fmt.Println("         a、这只是工程管理需要，便于识别一组函数，非技术实现需求")
    fmt.Println("         b、结果上使得go的状态和行为发生了分离,这是很多语言试图通过设计原则达成的目标")
    fmt.Println("         c、从实现上彻底摒弃了继承、多态这种复杂的OO机制，这些在主流OO社区也开始被弱化")
    fmt.Println("     3、方法传参上都是传值，这个值或者为基本类型或者为指针（或者叫引用）")
    fmt.Println("         a、这种传值方式，如果遇到结构体就是在栈空间复制原结构体内容，遇到引用（指针）类型，就是在栈空间存储引用")
    fmt.Println("         b、这种转换是编译器根据方法声明对变量类型进行自动转换的，程序员不需要关注，无需程序员手动转换")
    fmt.Println("         c、系统默认将数据拷贝到栈空间，避免方法的副作用,保证方法的内聚性，在读和写上都不依赖外部数据")
    fmt.Println("         d、如果比较在意这种拷贝带来的性能问题，go支持通过指针类型传递地址来修改")
    fmt.Println("         e、变量的本质是内存地址，而类型主要在于决定怎么看待使用这块内存里的数据")

    root.Travel()
}
