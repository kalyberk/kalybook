package orderbook

type OrderBook struct {
	Bids    map[float64]*Limit
	Asks    map[float64]*Limit
	BestBid *Order // Highest bid
	BestAsk *Order // Lowest ask
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Bids: make(map[float64]*Limit),
		Asks: make(map[float64]*Limit),
	}
}

func (b *OrderBook) AddOrder(order *Order) {
	var book map[float64]*Limit
	if order.BidOrAsk {
		book = b.Bids
	} else {
		book = b.Asks
	}

	limit, ok := book[order.Price]
	if !ok {
		limit = NewLimit(order.Price)
		book[order.Price] = limit
	}
	limit.Enqueue(order)

	if order.BidOrAsk && (b.BestBid == nil || order.Price > b.BestBid.Price) {
		b.BestBid = order
	} else if !order.BidOrAsk && (b.BestAsk == nil || order.Price < b.BestAsk.Price) {
		b.BestAsk = order
	}
}

func (b *OrderBook) CancelOrder(order *Order) {
	limit := order.Limit
	limit.Remove(order)
	if limit.IsEmpty() {
		b.DeleteLimit(limit)
	}
}

func (b *OrderBook) ExecOrder(order *Order) {
	for order.Remaining > 0 {
		matchingOrder := b.BestMatchingOrder(order)
		if matchingOrder == nil {
			break
		}
		matchingPrice := matchingOrder.Price
		matchingAmount := min(order.Remaining, matchingOrder.Remaining)

		order.Remaining -= matchingAmount
		matchingOrder.Remaining -= matchingAmount
		println("Matched", matchingAmount, "at", matchingPrice)

		if matchingOrder.Remaining == 0 {
			limit := matchingOrder.Limit
			limit.Dequeue()
			if limit.IsEmpty() {
				b.DeleteLimit(limit)
			}
			b.UpdateBestPrice(matchingOrder.BidOrAsk)
		}
	}

	if order.Remaining > 0 {
		b.AddOrder(order)
	}
}

// Find the best matching order for the given order
func (b *OrderBook) BestMatchingOrder(order *Order) *Order {
	var oppositeOrder *Order
	if order.BidOrAsk {
		oppositeOrder = b.BestAsk
	} else {
		oppositeOrder = b.BestBid
	}
	if oppositeOrder != nil && ((order.BidOrAsk && order.Price >= oppositeOrder.Price) || (!order.BidOrAsk && order.Price <= oppositeOrder.Price)) {
		return oppositeOrder
	}
	return nil
}

func (b *OrderBook) DeleteLimit(limit *Limit) {
	// TODO: What if limit is not empty?
	// TODO  What if limit is not in the book?
	var book map[float64]*Limit
	if limit.Orders.head.BidOrAsk {
		book = b.Bids
	} else {
		book = b.Asks
	}
	delete(book, limit.Price)
}

func (b *OrderBook) UpdateBestPrice(bidOrAsk bool) {
	if bidOrAsk {
		b.UpdateBestBid()
	} else {
		b.UpdateBestAsk()
	}
}

func (b *OrderBook) UpdateBestBid() {
	if len(b.Bids) > 0 {
		b.BestBid = b.Bids[maxPrice(b.Bids)].Orders.head
	} else {
		b.BestBid = nil
	}
}

func (b *OrderBook) UpdateBestAsk() {
	if len(b.Asks) > 0 {
		b.BestAsk = b.Asks[minPrice(b.Asks)].Orders.head
	} else {
		b.BestAsk = nil
	}
}

// TODO: maxPrice & minPrice implementations are not good enough
// If Bids and Asks would be stored in any other efficient data structure (e.g. heap, red-black tree, etc.)
// we could get the best price in O(1) time

func maxPrice(limits map[float64]*Limit) float64 {
	var maxPrice float64
	for price := range limits {
		if price > maxPrice {
			maxPrice = price
		}
	}
	return maxPrice
}

func minPrice(limits map[float64]*Limit) float64 {
	var minPrice float64
	for price := range limits {
		if minPrice == 0 || price < minPrice {
			minPrice = price
		}
	}
	return minPrice
}

func (b *OrderBook) BidCount() int {
	var count int
	for _, limit := range b.Bids {
		count += limit.Orders.len()
	}
	return count
}

func (b *OrderBook) AskCount() int {
	var count int
	for _, limit := range b.Asks {
		count += limit.Orders.len()
	}
	return count
}
