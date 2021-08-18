package queue

type IntQueue []int

func (queue *IntQueue) Enqueue(i int) {
	*queue = append(*queue, i)
}

func (queue *IntQueue) Dequeue() int {
	result := (*queue)[0]
	*queue = (*queue)[1:]
	return result
}
