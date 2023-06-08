package file

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log/logger"
)

type writer struct {
	io.Writer
	config config
}

func New(registry registry.Registry) logger.Writer {
	w := &writer{}
	if err := registry.Unmarshal(&w.config); err != nil {
		panic(err)
	}
	if !w.config.Enabled {
		return nil
	}
	w.config.initialize()

	name := w.config.Name
	now := time.Now().UTC()
	name = strings.Replace(name, "%year", fmt.Sprintf("%04d", now.Year()), -1)
	name = strings.Replace(name, "%month", fmt.Sprintf("%02d", now.Month()), -1)
	name = strings.Replace(name, "%day", fmt.Sprintf("%02d", now.Day()), -1)
	name = strings.Replace(name, "%hour", fmt.Sprintf("%02d", now.Hour()), -1)
	name = strings.Replace(name, "%minute", fmt.Sprintf("%02d", now.Minute()), -1)
	name = strings.Replace(name, "%second", fmt.Sprintf("%02d", now.Second()), -1)

	path := path.Join(w.config.Dir, name)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	w.Writer = f
	return w
}

func (w *writer) Start() {

}

func (w *writer) Level() logger.Level {
	return w.config.Level
}
