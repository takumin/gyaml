package config

type Option interface {
	Apply(*Config)
}

type LogLevel string

func (o LogLevel) Apply(c *Config) {
	c.LogLevel = string(o)
}

type Path string

func (o Path) Apply(c *Config) {
	c.Paths = append(c.Paths, string(o))
}

type Includes string

func (o Includes) Apply(c *Config) {
	c.Extention.Includes = string(o)
}

type Excludes string

func (o Excludes) Apply(c *Config) {
	c.Extention.Excludes = string(o)
}
