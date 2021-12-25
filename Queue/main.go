package main

import "fmt"

//struct represent queue that holds slice
type Queue struct {
	List []int
}

//add new value in queue
func (q *Queue) Push(i int)  {
	q.List = append(q.List, i)
}

//remove first value and returns removed value
func (q *Queue) Pop() int {
	firstValue := q.List[0]
	q.List = q.List[1:]
	return firstValue
}

func main()  {
	//new copy of queue
	q := Queue{}

	//display empty queue
	fmt.Println(q)

	//adding values to queue
	q.Push(111)
	q.Push(222)
	q.Push(333)
	q.Push(444)

	fmt.Println(q)

	//remove and display removed values
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println(q)
}
