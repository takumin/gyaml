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

type Include string

func (o Include) Apply(c *Config) {
	c.Extention.Includes = append(c.Extention.Includes, string(o))
}

type Exclude string

func (o Exclude) Apply(c *Config) {
	c.Extention.Excludes = append(c.Extention.Excludes, string(o))
}
