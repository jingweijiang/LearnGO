package main

import "fmt"

type SNode2 struct {
	Node // 默认字段名同类型名，即`Node`
}

func (s *SNode2) PostOrder() {
	if s == nil {
		return
	}

	s.Left.PreOrder()    // 本质上是：s.Node.Left.PreOrder()
	fmt.Println(s.Value) // 本质上是：s.Node.Value
	s.Right.PreOrder()   // 本质上是：s.Node.Right.PreOrder
}

// 实际上编译器默认会给`*SNode`自动生成内嵌类型`Node`的同名方法，
// 但是这里有空指针的风险，因此我们没有采用默认实现
// func (s *SNode) PreOrder() {
//   s.Node.PreOrder()
//}

func (s *SNode2) PreOrder() {
	if s != nil {
		s.Node.PreOrder()
	}
}
