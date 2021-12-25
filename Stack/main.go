package main

import "fmt"

//Stack data structure
type Stack struct {
	List []int
}

//add new value in stack
func (s *Stack) Push(i int) {
	s.List = append(s.List, i)
}

//remove last value and returns removed index and value
func (s *Stack) Pop() (int, int) {
	lastIndex := len(s.List)-1
	lastValue := s.List[lastIndex]
	s.List = s.List[:lastIndex]
	return lastIndex, lastValue
}

func main()  {
	//new copy of structure
	s := Stack{}

	//displaying copy of structure
	fmt.Println(s)

	//adding values
	s.Push(111)
	s.Push(222)
	s.Push(333)
	s.Push(444)

	fmt.Println(s)

	//remove and display values
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	fmt.Println(s)
}
