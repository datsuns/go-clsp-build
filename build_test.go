package main

import (
	"fmt"
	"testing"
)

func TestBuild_load(t *testing.T) {
	c := LspConfig{
		Location: "./tests",
		Patterns: []string{"c", "cpp"},
		Targets:  []string{""},
		Flags:    []string{},
	}
	e, err := buildStructure(&c, []string{})
	if err != nil {
		t.Errorf("build error [%v]\n", e)
	}
	if len(*e) == 0 {
		t.Errorf("build size error\n")
	}
	fmt.Println(*e)

}
