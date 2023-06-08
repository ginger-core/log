package log

func (l *_logger) Start() {
	writers := l.GetWriters()
	for _, w := range writers {
		w.Start()
	}
}
