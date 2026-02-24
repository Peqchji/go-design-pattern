package marketplace

type StockUpdateEvent struct {
	ProductID string
	IsInStock bool
	Price     float64
}
