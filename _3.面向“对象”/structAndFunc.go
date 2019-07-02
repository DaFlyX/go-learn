package main

import (
	"fmt"
)

/*
	go语言仅支持封装，不支持继承和多态


*/
//定义结构体
type treeNode struct {
	value       int
	left, right *treeNode
}

//工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}
func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.print()
	root.setValue(100)
	pRoot := &root
	pRoot.print()

	fmt.Println()
}
