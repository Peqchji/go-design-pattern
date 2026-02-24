package observer

type Observer[E any] interface {
	ID() string
	OnNotify(event E)
}

type Subject[E any] interface {
	Register(observer Observer[E])
	Deregister(observer Observer[E])
	Notify(event E)
}
