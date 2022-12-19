package config

import "strconv"

type InMemoryConfigReader struct {
	values map[string]string
}

func NewInMemoryConfigReader() *InMemoryConfigReader {
	return &InMemoryConfigReader{
		values: make(map[string]string),
	}
}

func (c *InMemoryConfigReader) SetValue(name string, value string) {
	c.values[name] = value
}

func (c *InMemoryConfigReader) String(name string, def string) (string, bool) {
	v, exst := c.values[name]

	if !exst {
		return def, false
	}

	return v, true
}
func (c *InMemoryConfigReader) Int32(name string, def int) (int, bool) {
	v, exst := c.values[name]

	if !exst {
		return def, false
	}

	ival, err := strconv.Atoi(v)

	if err != nil {
		return -1, false
	}

	return ival, true
}

func (c *InMemoryConfigReader) Bool(name string, def bool) (bool, bool) {
	v, exst := c.values[name]

	if !exst {
		return false, false
	}

	bval, err := strconv.ParseBool(v)

	if err != nil {
		return def, false
	}

	return bval, true
}
