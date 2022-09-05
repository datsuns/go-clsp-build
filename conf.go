package main

import (
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type LspConfig struct {
	Location string   `yaml:"location"`
	Patterns []string `yaml:"patterns"`
	Targets  []string `yaml:"targets"`
	Flags    []string `yaml:"flags"`
}

type LspConfigTableFixed struct {
	Packages    []string `yaml:"Packages"`
	CommonFlags []string `yaml:"CommonFlags"`
}

type LspConfigTable struct {
	LspConfigTableFixed
	Entries map[string]LspConfig
}

type ClangdOptionEntry struct {
	Directory string   `json:"directory"`
	Arguments []string `json:"arguments"`
	File      string   `json:"file"`
}

func (t *LspConfigTable) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var params LspConfigTableFixed
	if err := unmarshal(&params); err != nil {
		return err
	}
	var entries map[string]LspConfig
	if err := unmarshal(&entries); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	t.Packages = params.Packages
	t.CommonFlags = params.CommonFlags
	t.Entries = entries
	return nil
}

func unmarshalConf(raw []byte) (*LspConfigTable, error) {
	ret := &LspConfigTable{}
	err := yaml.Unmarshal(raw, ret)
	return ret, err
}

func loadConfFromIo(r io.Reader) (*LspConfigTable, error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return unmarshalConf(raw)
}

func loadConfFromFile(path string) (*LspConfigTable, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return unmarshalConf(raw)
}
