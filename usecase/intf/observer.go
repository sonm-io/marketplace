package intf

type Publisher interface {
	Publish(value interface{})
}

type Observable []EventHandler

func (observers *Observable) AddObserver(a EventHandler) {
	*observers = append(*observers, a)
}

func (observers Observable) Publish(event Event) {
	for _, obs := range observers {
		obs.Handle(event)
	}
}
