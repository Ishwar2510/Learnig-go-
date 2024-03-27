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
	data, err := q.queue.RemoveFirst()
	return data, err
}
func (q *Queue) PrintQueue() {
	q.queue.PrintList()
}
