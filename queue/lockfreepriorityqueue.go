package lockfreepriorityqueue

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

var ErrPriorityExisted = errors.New("error: priority is existed")

type LockFreePriorityQueue struct {
	next unsafe.Pointer
}

type node struct {
	priority uint64
	value    interface{}
	next     unsafe.Pointer
}

func NewLKQueue() *LockFreePriorityQueue {
	n := unsafe.Pointer(&node{})
	lk := &LockFreePriorityQueue{next: n}
	return lk
}

func (q *LockFreePriorityQueue) Push(v interface{}, priority uint64) error {
	n := &node{
		priority: priority,
		value:    v}
	head := Load(&q.next)
	for {

		if head.value == nil {
			cas(&q.next, head, n)
			return nil
		}
		if head.priority > priority {
			if cas(&n.next, nil, head) && cas(&q.next, head, n) {
				return nil
			}
		} else if head.priority == priority {
			return ErrPriorityExisted
		} else if head.priority < priority {
			if head.next == nil {
				if cas(&head.next, nil, n) {
					return nil
				}
			}
			head = Load(&head.next)
		}
	}
}

func (q *LockFreePriorityQueue) Pop() interface{} {

	head := Load(&q.next)
	if head.value == nil {
		return nil
	} else {
		v := head.value
		tailNext := Load(&head.next)
		if cas(&q.next, head, tailNext) {
			return v
		}
	}
	return nil
}

func Load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}
