package orderbook

import (
	"math/rand"
	"testing"
)

func TestOrdersQueue(t *testing.T) {
	t.Run("enqueue", testOrdersQueueEnqueue)
	t.Run("dequeue", testOrdersQueueDequeue)
}

func testOrdersQueueEnqueue(t *testing.T) {
	oq := newOrdersQueue()
	n := 100

	for i := 0; i < n; i++ {
		oq.enqueue(&Order{
			ID:       i,
			Amount:   rand.Float64() * 100,
			BidOrAsk: rand.Intn(2) == 0,
		})
	}
	if oq.len() != n {
		t.Errorf("Expected length %d, got %d", n, oq.len())
	}
}

func testOrdersQueueDequeue(t *testing.T) {
	oq := newOrdersQueue()
	n := 100

	for i := 0; i < n; i++ {
		oq.enqueue(&Order{
			ID:       i,
			Amount:   rand.Float64() * 100,
			BidOrAsk: rand.Intn(2) == 0,
		})
	}
	for i := 0; i < n; i++ {
		oq.dequeue()
	}
	if oq.len() != 0 {
		t.Errorf("Expected length %d, got %d", 0, oq.len())
	}
}
