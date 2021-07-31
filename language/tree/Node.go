package tree

import "fmt"

type TreeNode struct {
    value       int
    left, right *TreeNode
}

func (node *TreeNode) travel() {
    if node == nil {
        return
    }
    fmt.Println(node.value)
    // 注意这里不需要判断node.left != nil
    node.left.travel()
    node.right.travel()
}

func (node *TreeNode) setValue(v int) {
    node.value = v
}

func (node TreeNode) setValue1(v int) {
    node.value = v
}

func (node TreeNode) printValue2() {
    fmt.Println(node.value)
}

// 创建工厂
func createTreeNode(value int) *TreeNode {
    // 注意这里返回了局部变量的地址，是不是很神奇
    return &TreeNode{value: value}
}
