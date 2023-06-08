package file

func (w *writer) Write(p []byte) (n int, err error) {
	if !w.config.Enabled {
		return 0, nil
	}
	n, err = w.Writer.Write(p)
	// w.Writer.Write([]byte{'\n'})
	return n, err
}
