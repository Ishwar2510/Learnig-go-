package main

type Queue struct {
	queue *Linkedlist
}

func NewQueue() *Queue {
	return &Queue{queue: NewLinkedList()}

}
func (q *Queue) Add(data interface{}) error {
	err := q.queue.Insert(data)
	return err
}
func (q *Queue) Remove() (interface{}, error) {
	data, _ := q.queue.RemoveFirst()
	return data, nil
}
func (q *Queue) Poll() (interface{}, error) {
	data, _ := q.queue.RemoveFirst()
	return data, nil
}
func (q *Queue) IsEmpty() bool {
	if err := q.queue.isEmpty(); err != nil {
		return true
	}
	return false
}
func (q *Queue) Size() int {
	return q.queue.Size()
}
func (q *Queue) Peek() (interface{}, error) {
	if data, err := q.queue.PeekFirst(); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
func (q *Queue) PrintQueue() {
	q.queue.PrintList()
}
