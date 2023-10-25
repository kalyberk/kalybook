package orderbook


type Order struct {
	ID        int
	BidOrAsk  bool // true for a bid and false for an ask
	Price     float64
	Amount    float64
	Filled    float64
	Remaining float64
	Limit *Limit
	Prev  *Order
	Next  *Order
}

func NewOrder(id int, bidOrAsk bool, price float64, amount float64) *Order {
	return &Order{
		ID:        id,
		BidOrAsk:  bidOrAsk,
		Price:     price,
		Amount:    amount,
		Filled:    0,
		Remaining: amount,
	}
}
