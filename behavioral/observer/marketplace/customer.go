package marketplace

import (
	"design_pattern/behavioral/observer"
	"fmt"
)

var _ observer.Observer[StockUpdateEvent] = &Customer{}

type Customer struct {
	id string
	lastUpdate string
}

func NewCustomer(id string) *Customer {
	return &Customer{
		id: id,
	}
}

func (c *Customer) ID() string {
	return c.id
}

func (c *Customer) OnNotify(event StockUpdateEvent) {
	c.lastUpdate = fmt.Sprintf(
		"Customer <%s> received update: Product <%s> is now <%v>", 
		c.id, 
		event.ProductID, 
		event.IsInStock,
	)
}

func (c *Customer) LastUpdate() string {
	return c.lastUpdate
}
