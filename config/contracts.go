package config

type IConfigReader interface {
	String(name string, def string) (string, bool)
	Int32(name string, def int) (int, bool)
	Bool(name string, def bool) (bool, bool)
}
