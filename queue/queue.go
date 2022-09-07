package queue

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dqueue() int {
	itemToRemove := q.items[0]
	q.items = q.items[1:]
	return itemToRemove
}
