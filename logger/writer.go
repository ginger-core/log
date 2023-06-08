package logger

import (
	"io"
)

type Writer interface {
	io.Writer
	Start()
	Level() Level
}
