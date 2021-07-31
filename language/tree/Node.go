package tree

import "fmt"

type Node struct {
    Value       int
    Left, Right *Node
}

func (node *Node) Travel() {
    if node == nil {
        return
    }
    fmt.Println(node.Value)
    // 注意这里不需要判断node.left != nil
    node.Left.Travel()
    node.Right.Travel()
}

func (node *Node) SetValue(v int) {
    node.Value = v
}

func (node Node) SetValue1(v int) {
    node.Value = v
}

func (node Node) PrintValue2() {
    fmt.Println(node.Value)
}

// 创建工厂
func CreateTreeNode(value int) *Node {
    // 注意这里返回了局部变量的地址，是不是很神奇
    return &Node{Value: value}
}
