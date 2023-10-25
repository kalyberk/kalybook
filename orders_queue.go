package orderbook

// A doubly linked list for orders (FIFO) at a certain price level
type ordersQueue struct {
	head   *Order
	tail   *Order
	length int
}

func newOrdersQueue() *ordersQueue {
	return &ordersQueue{}
}

func (oq *ordersQueue) len() int {
	return oq.length
}

func (oq *ordersQueue) enqueue(order *Order) {
	if oq.length == 0 {
		oq.head = order
		oq.tail = order
	} else {
		oq.tail.Next = order
		order.Prev = oq.tail
		oq.tail = order
	}
	oq.length++
}

func (oq *ordersQueue) dequeue() *Order {
	if oq.length == 0 {
		return nil
	}
	order := oq.head
	oq.head = order.Next
	oq.length--
	return order
}

func (oq *ordersQueue) remove(order *Order) bool {
	if oq.length == 0 {
		return false
	}
	prev := order.Prev
	next := order.Next
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
	order.Prev = nil
	order.Next = nil
	if order == oq.head {
		oq.head = next
	}
	if order == oq.tail {
		oq.tail = prev
	}
	oq.length--
	return true
}
