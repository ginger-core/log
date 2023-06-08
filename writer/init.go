package writer

import (
	"fmt"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-core/log/writer/file"
	"github.com/ginger-core/log/writer/stdout"
)

func New(w string, registry registry.Registry) logger.Writer {
	switch w {
	case "stdout":
		return stdout.New(registry)
	case "file":
		return file.New(registry)
	}
	panic(fmt.Sprintf("writer %s not found", w))
}
