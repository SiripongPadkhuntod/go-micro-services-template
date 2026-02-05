package property

import (
	"os"
	"time"
)

type Config struct {
	ServiceName string
	Port        string
}

// Deadline implements [context.Context].
func (c *Config) Deadline() (deadline time.Time, ok bool) {
	panic("unimplemented")
}

// Done implements [context.Context].
func (c *Config) Done() <-chan struct{} {
	panic("unimplemented")
}

// Err implements [context.Context].
func (c *Config) Err() error {
	panic("unimplemented")
}

// Value implements [context.Context].
func (c *Config) Value(key any) any {
	panic("unimplemented")
}

func Load() *Config {
	return &Config{
		ServiceName: getenv("SERVICE_NAME", "template-service"),
		Port:        getenv("PORT", "8080"),
	}
}

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
