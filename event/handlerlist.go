package event

type HandlerList []EventHandler

func NewHandlerList() *HandlerList {
	return &HandlerList{}
}

func (ll *HandlerList) Add(h EventHandler) {
	*ll = append(*ll, h)
}

func (ll *HandlerList) ParallelNotify(evt *Event) {
	for _, h := range *ll {
		go h(evt)
	}
}
