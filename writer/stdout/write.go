package stdout

func (w *writer) Write(p []byte) (n int, err error) {
	n, err = w.Writer.Write(p)
	w.Writer.Write([]byte{'\n'})
	return n, err
}
