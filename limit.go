package orderbook

type Limit struct {
	Price       float64
	Orders      *ordersQueue
	TotalAmount float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: newOrdersQueue(),
	}
}

func (l *Limit) IsEmpty() bool {
	return l.Orders.len() == 0
}

func (l *Limit) Enqueue(order *Order) {
	l.Orders.enqueue(order)
	l.TotalAmount += order.Remaining
	order.Limit = l
}

func (l *Limit) Dequeue() *Order {
	order := l.Orders.dequeue()
	if order != nil {
		l.TotalAmount -= order.Remaining
		order.Limit = nil
	}
	return order
}

func (l *Limit) Remove(order *Order) {
	if order.Limit != l {
		return
	}
	if l.Orders.remove(order) {
		l.TotalAmount -= order.Remaining
		order.Limit = nil
	}
}
