package logger

type Config struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level Level
	// Encoding sets the logger's encoding. Valid values are "json" and
	// "console", as well as any third-party encodings registered via
	// RegisterEncoder.
	Encoding string
	//
	Writers map[string]any
}

func (c *Config) initialize() {
	if c.Encoding == "" {
		c.Encoding = "console"
	}
}
