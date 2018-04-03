package queue

import (
	"sync"
)

type Queue struct {
	items []QueueItem
	mutex *sync.Mutex
}

type QueueItem interface {
	GetId() int
}

func New() *Queue {
	return &Queue{mutex: &sync.Mutex{}}
}

func (q *Queue) Push(obj QueueItem) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, obj)
}

func (q *Queue) Pop() QueueItem {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:len(q.items)]

	return item
}

func (q *Queue) RemoveById(id int) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	index := q.findIndex(id)
	if index == -1 {
		return false
	}

	q.items = append(q.items[0:index], q.items[index+1:]...)
	return true
}

func (q *Queue) Remove(obj QueueItem) bool {
	return q.RemoveById(obj.GetId())
}

func (q *Queue) IsEmpty() bool {
	return q.Count() == 0
}

func (q *Queue) Count() int {
	return len(q.items)
}

func (q *Queue) Exist(obj QueueItem) bool {
	return q.ExistById(obj.GetId())
}

func (q *Queue) ExistById(id int) bool {
	return q.findIndex(id) > -1
}

func (q *Queue) findIndex(id int) int {
	for i := 0; i < len(q.items); i++ {
		if q.items[i].GetId() == id {
			return i
		}
	}

	return -1
}
