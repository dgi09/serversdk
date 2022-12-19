package event

type EventHandler = func(evt *Event)

type IEventEmitter interface {
	Emit(name string, args ...any)
	Subscribe(name string, handler EventHandler)
}
