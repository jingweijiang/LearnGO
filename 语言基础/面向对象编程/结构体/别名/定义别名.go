package main

import "fmt"

type Stack []int

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	data := *s
	v := data[len(data)-1]
	*s = data[:len(data)-1]
	return v
}

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop(), s.Pop(), s.Pop()) // 3 2 1
}
