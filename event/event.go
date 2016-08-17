package event

// Event is any event, may contain a string shell command OR an error
type Event struct {
	Shell string
	Error error
}

// Error returns an error-kind of event
func Error(err error) Event {
	return Event{Error: err}
}

// Shell returns a shell-kind of event
func Shell(shell string) Event {
	return Event{Shell: shell}
}
