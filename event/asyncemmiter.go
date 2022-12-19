package event

type AsyncEventEmmiter struct {
	handlers map[string]*HandlerList
	context  any
}

func NewAsyncEventEmmiter(context any) *AsyncEventEmmiter {
	return &AsyncEventEmmiter{
		handlers: make(map[string]*HandlerList),
		context:  context,
	}
}

func (e *AsyncEventEmmiter) Emit(name string, args ...any) {
	if list, found := e.handlers[name]; found {
		evt := NewEvent(args, e.context)
		list.ParallelNotify(evt)
	}
}

func (e *AsyncEventEmmiter) Subscribe(name string, handler EventHandler) {
	list, found := e.handlers[name]

	if !found {
		list = NewHandlerList()
		e.handlers[name] = list
	}

	list.Add(handler)
}
