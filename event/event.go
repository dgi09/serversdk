package event

type Event struct {
	args    []any
	context any
}

func NewEvent(args []any, context any) *Event {
	return &Event{
		args:    args,
		context: context,
	}
}

func (e *Event) String(idx int) string {
	if idx < 0 || idx >= len(e.args) {
		panic("Arg index out of bounds")
	}

	return e.args[idx].(string)
}

func (e *Event) Int(idx int) int {
	if idx < 0 || idx >= len(e.args) {
		panic("Arg index out of bounds")
	}

	return e.args[idx].(int)
}

func (e *Event) Int64(idx int) int64 {
	if idx < 0 || idx >= len(e.args) {
		panic("Arg index out of bounds")
	}

	return e.args[idx].(int64)
}

func (e *Event) Uint64(idx int) uint64 {
	if idx < 0 || idx >= len(e.args) {
		panic("Arg index out of bounds")
	}

	return e.args[idx].(uint64)
}

func (e *Event) Bool(idx int) bool {
	if idx < 0 || idx >= len(e.args) {
		panic("Arg index out of bounds")
	}

	return e.args[idx].(bool)
}

func (e *Event) Context() any {
	return e.context
}
