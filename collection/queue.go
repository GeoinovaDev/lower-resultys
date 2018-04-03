package collection

type Queue struct {
	items []interface{}
}

func New1() *Queue {
	return &Queue{}
}

func (q *Queue) Push(obj interface{}) {
	q.items = append(q.items, obj)
}

func (q *Queue) Pop(obj interface{}) interface{} {
	return q.items[len(q.items)-1]
}

func (q *Queue) Remove(obj interface{}) {

}
