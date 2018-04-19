package queue

import (
	"sync"
)

// Queue é a estrutura contendo os itens da fila e a sincronização
type Queue struct {
	items []Item
	mutex *sync.Mutex
}

// New cria uma nova fila
func New() *Queue {
	return &Queue{mutex: &sync.Mutex{}}
}

// Push adiciona item no inicio da fila
func (q *Queue) Push(obj Item) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, obj)
}

// Pop retorna o primeiro item adicionado
func (q *Queue) Pop() Item {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:len(q.items)]

	return item
}

// RemoveByID remove item pelo seu id
// Retorna verdadeiro se o item esta na fila
func (q *Queue) RemoveByID(id int) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	index := q.findIndex(id)
	if index == -1 {
		return false
	}

	q.items = append(q.items[0:index], q.items[index+1:]...)
	return true
}

// Remove exclui um item da fila
// Retorna verdadeiro se houve sucesso
func (q *Queue) Remove(obj Item) bool {
	return q.RemoveByID(obj.GetID())
}

// IsEmpty retorna se a fila esta vazia
func (q *Queue) IsEmpty() bool {
	return q.Count() == 0
}

// Count retorna o total de itens na fila
func (q *Queue) Count() int {
	return len(q.items)
}

// Exist retorna se um item existe na fila
func (q *Queue) Exist(obj Item) bool {
	return q.ExistByID(obj.GetID())
}

// ExistByID retorna se um item existe na fila pelo seu id
func (q *Queue) ExistByID(id int) bool {
	return q.findIndex(id) > -1
}

func (q *Queue) findIndex(id int) int {
	for i := 0; i < len(q.items); i++ {
		if q.items[i].GetID() == id {
			return i
		}
	}

	return -1
}
