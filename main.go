package main

import "fmt"

//Stack data structure
type Stack struct {
	List []int
}

func (s *Stack) Push(i int) {
	s.List = append(s.List, i)
}

func (s *Stack) Pop() int {
	last := s.List[len(s.List)-1]
	s.List = s.List[:len(s.List)-1]
	return last
}

func main()  {
	s := Stack{}
	fmt.Println(s)
	s.Push(111)
	s.Push(222)
	s.Push(333)
	s.Push(444)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s)
}
