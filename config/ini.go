package config

import (
	"gopkg.in/ini.v1"
)

type IniConfigReader struct {
	fileName    string
	inMemReader *InMemoryConfigReader
}

func NewIniConfigReader(fileName string) *IniConfigReader {
	return &IniConfigReader{
		fileName: fileName,
	}
}

func (r *IniConfigReader) Load() error {
	cfg, err := ini.Load(r.fileName)
	if err != nil {
		return err
	}

	r.inMemReader = NewInMemoryConfigReader()

	sections := cfg.Sections()

	for _, s := range sections {
		sname := s.Name()
		keys := s.Keys()

		for _, k := range keys {
			name := sname + "_" + k.Name()

			r.inMemReader.SetValue(name, k.Value())
		}
	}

	return nil
}

func (r *IniConfigReader) String(name string, def string) (string, bool) {
	return r.inMemReader.String(name, def)
}

func (r *IniConfigReader) Int32(name string, def int) (int, bool) {
	return r.inMemReader.Int32(name, def)
}

func (r *IniConfigReader) Bool(name string, def bool) (bool, bool) {
	return r.inMemReader.Bool(name, def)
}
