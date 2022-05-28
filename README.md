# lock_free_priority_queue
Lock-free priority queue, using CAS to ensure concurrency security

# test
```
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
```

result:

```
value_1
value_2
value_3
value_4
value_5
value_6
value_7
value_8
value_9
value_10
value_11
value_12
value_13
value_14
value_15
value_16
value_17
value_18
value_19
value_20
```
