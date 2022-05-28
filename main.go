package main

import (
	"fmt"
	"lock_free_priority_queue/queue"
	"strconv"
)

func main() {

	q := lockfreepriorityqueue.NewLKQueue()
	for i := 10; i >= 1; i-- {
		q.Push("value_"+strconv.Itoa(i), uint64(i))
	}
	for i := 9; i <= 20; i++ {
		q.Push("value_"+strconv.Itoa(i), uint64(i))
	}
	for i := 1; i <= 20; i++ {
		fmt.Println(q.Pop())
	}
}
