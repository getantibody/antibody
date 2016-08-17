package event

// Event is any event, may contain a string shell command OR an error
type Event struct {
	Shell string
	Error error
}

func Error(err error) Event {
	return Event{Error: err}
}

func Shell(shell string) Event {
	return Event{Shell: shell}
}
