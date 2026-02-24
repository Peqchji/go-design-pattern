package marketplace

import (
	"design_pattern/behavioral/observer"
)


type Product struct {
	ID        string
	Name      string
	inStock   bool
	price     float64
	observers map[string]observer.Observer[StockUpdateEvent]
}

func NewProduct(id string, name string, price float64) Product {
	return Product{
		ID:    id,
		Name:  name,
		price: price,
		observers: make(map[string]observer.Observer[StockUpdateEvent]),
	}
}

func (p *Product) Register(observer observer.Observer[StockUpdateEvent]) {
	oid := observer.ID()
	if _, exists := p.observers[oid]; exists {
		return
	}

	p.observers[oid] = observer
}

func (p *Product) Deregister(observer observer.Observer[StockUpdateEvent]) {
	delete(p.observers, observer.ID())
}

func (p *Product) Notify(event StockUpdateEvent) {
	for _, observer := range p.observers {
		observer.OnNotify(event)
	}
}

func (p *Product) UpdateStock(inStock bool) {
	p.inStock = inStock
}

func (p *Product) UpdatePrice(newPrice float64) {
	p.price = newPrice
}
