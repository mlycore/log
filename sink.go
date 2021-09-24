package log

import "bytes"

type logEvent struct {
	fields Fields
	ctx    Context
}

// Sink defines a Sink
type Sink interface {
	Receiver() chan logEvent
	Output() error
}

type StdioSink struct {
	buf bytes.Buffer
}

func (s *StdioSink) Receiver() chan logEvent {
	return make(chan logEvent)
}

func (s *StdioSink) Output() error {
	return nil
}
