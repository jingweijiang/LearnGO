package main

import "fmt"

type Node struct {
	Value int   `json:"value"`
	Left  *Node `json:"left"`  // left child
	Right *Node `json:"right"` // right child
}

// Node 判断是否叶子节点
func (n Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

// PreOrder 前缀遍历
func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	n.Left.PreOrder()
	fmt.Println(n.Value)
	n.Right.PreOrder()
}

type SNode struct {
	node Node
}

func (s *SNode) Postorder(value int) {
	if s == nil {
		return
	}

	s.node.Left.PreOrder()
	fmt.Println(s.node.Value)
	s.node.Right.PreOrder()
}

func (s *SNode) PreOrder() {
	if s != nil {
		s.node.PreOrder()
	}
}

func main() {
	n := Node{
		Value: 1,
		Left:  &Node{Value: 2, Left: nil, Right: nil},
		Right: &Node{Value: 3, Left: nil, Right: nil}}

	fmt.Println(n)
}
