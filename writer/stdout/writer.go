package stdout

import (
	"io"
	"os"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log/logger"
)

type writer struct {
	io.Writer
	config config
}

func New(registry registry.Registry) logger.Writer {
	w := &writer{
		Writer: os.Stdout,
	}
	if err := registry.Unmarshal(&w.config); err != nil {
		panic(err)
	}
	if !w.config.Enabled {
		return nil
	}
	return w
}

func (w *writer) Start() {

}

func (w *writer) Level() logger.Level {
	return w.config.Level
}
