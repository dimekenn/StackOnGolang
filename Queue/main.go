package main

import "fmt"

type Queue struct {
	List []int
}

func (q *Queue) Push(i int)  {
	q.List = append(q.List, i)
}

func (q *Queue) Pop() (int, int) {
	lastIndex := len(q.List) -1
	lastValue := q.List[lastIndex]
	q.List = q.List[1:]
	return lastIndex, lastValue
}

func main()  {
	q := Queue{}

	fmt.Println(q)

	q.Push(111)
	q.Push(222)
	q.Push(333)
	q.Push(444)

	fmt.Println(q)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println(q)
}
