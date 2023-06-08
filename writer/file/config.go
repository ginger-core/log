package file

import "github.com/ginger-core/log/logger"

type config struct {
	Enabled bool
	Dir     string
	Name    string
	Level   logger.Level
}

func (c *config) initialize() {
	if c.Dir == "" {
		panic("log directory must be set")
	}
	if c.Name == "" {
		c.Name = "%year-%month-%day-%hour:%minute:%second.log"
	}
}
