package observer_test

import (
	"design_pattern/behavioral/observer/marketplace"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ProductNotifyTestCase struct {
	name string
	customers []*marketplace.Customer
	event *marketplace.StockUpdateEvent
	product marketplace.Product
	newStockStatus bool
	expected []string
}

func TestProductNotify(t *testing.T) {
	table := []ProductNotifyTestCase{
		{
			name: "Should notify all customers when stock is updated",
			customers: []*marketplace.Customer{
				marketplace.NewCustomer("1"),
				marketplace.NewCustomer("2"),
				marketplace.NewCustomer("3"),
			},
			event: &marketplace.StockUpdateEvent{
				ProductID: "1",
				IsInStock: true,
				Price:     100,
			},
			product: marketplace.NewProduct("1", "Product 1", 100),
			newStockStatus: true,
			expected: []string{
				"Customer <1> received update: Product <1> is now <true>",
				"Customer <2> received update: Product <1> is now <true>",
				"Customer <3> received update: Product <1> is now <true>",
			},
		},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			for _, customer := range tc.customers {
				tc.product.Register(customer)
			}
			
			tc.product.UpdateStock(tc.newStockStatus)
			if tc.event != nil {
				tc.product.Notify(*tc.event)
			}
			
			for i, customer := range tc.customers {
				if tc.event != nil {
					assert.Equal(t, tc.expected[i], customer.LastUpdate())
				} else {
					assert.Equal(t, "", customer.LastUpdate())
				}
			}
		})
	}
}