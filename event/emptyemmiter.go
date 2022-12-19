package event

type EmptyEventEmmiter struct{}

func NewEmptyEventEmmiter() *EmptyEventEmmiter {
	return &EmptyEventEmmiter{}
}

func (e *EmptyEventEmmiter) Emit(name string, args ...any)               {}
func (e *EmptyEventEmmiter) Subscribe(name string, handler EventHandler) {}
