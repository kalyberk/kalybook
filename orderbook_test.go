package orderbook

import (
	"testing"
)

func TestOrderBook(t *testing.T) {
	t.Run("NewOrderBook", testOrderBookNewOrderBook)
	t.Run("AddOrder", testOrderBookAddOrder)
	t.Run("CancelOrder", testOrderBookCancelOrder)
}

func testOrderBookNewOrderBook(t *testing.T) {
	ob := NewOrderBook()
	if ob == nil {
		t.Errorf("Expected OrderBook, got nil")
	}
}

func testOrderBookAddOrder(t *testing.T) {
	ob := NewOrderBook()
	n := 1000
	price := 100.0
	for i := 0; i < n; i++ {
		ob.AddOrder(NewOrder(i, true, price+float64(i), 1.0))
		ob.AddOrder(NewOrder(i, false, price-float64(i), 1.0))
	}
	bidCount := ob.BidCount()
	if bidCount != n {
		t.Errorf("Expected %d bids, got %d", n, bidCount)
	}
	askCount := ob.AskCount()
	if askCount != n {
		t.Errorf("Expected %d asks, got %d", n, askCount)
	}
}

func testOrderBookCancelOrder(t *testing.T) {
	ob := NewOrderBook()
	n := 5
	price := 100.0
	for i := 0; i < n; i++ {
		ob.AddOrder(NewOrder(i, true, price, 1.0))
	}
	bidCount := ob.BidCount()
	if bidCount != n {
		t.Errorf("Expected %d bids, got %d", n, bidCount)
	}
	// Cancel the third order
	thirdOrder := ob.Bids[price].Orders.head.Next.Next
	ob.CancelOrder(thirdOrder)
	bidCount = ob.BidCount()
	if bidCount != n-1 {
		t.Errorf("Expected %d bids, got %d", n-1, bidCount)
	}
}

