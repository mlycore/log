package log

import (
	"bytes"
	"time"
)

const flushInterval = 30 * time.Second

func (l *Logger) flushDaemon() {
	for _ = range time.NewTicker(flushInterval).C {
		l.lockAndFlushAll()
	}
}

func (l *Logger) lockAndFlushAll() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.flushAll()
}

func (l *Logger) flushAll() {
	ch := l.Sink.Receiver()
	var buf *bytes.Buffer
	for {
		select {
		case e := <-ch:
			{
				buf = bytes.NewBuffer([]byte(e.fields.Msg))
				l.Writer.Write(buf.Bytes())
			}
		default:
			{

			}
		}
	}
}
