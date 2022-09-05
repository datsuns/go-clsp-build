package main

import (
	"strings"
	"testing"
)

func TestConf_load(t *testing.T) {
	s := `
Packages:
  - p1
  - p2
CommonFlags:
  - cf1
  - cf2
  - cf3
p1:
  location:
    loc_p1
  patterns:
    - pat_p1_1
    - pat_p1_2
  targets:
    - t_p1_1
    - t_p1_2
  flags:
    - f_p1_1
    - f_p1_2
p2:
  location:
    loc_p2
  patterns:
    - pat_p2_1
    - pat_p2_2
  targets:
    - t_p2_1
    - t_p2_2
  flags:
    - f_p2_1
    - f_p2_2
   `
	c, e := loadConfFromIo(strings.NewReader(s))
	if e != nil {
		t.Errorf("load error [%v]", e)
	}
	if len(c.Packages) != 2 {
		t.Errorf("package size error 2 != [%v]", len(c.Packages))
	}
	if c.Packages[0] != "p1" {
		t.Errorf("package load error p1 != [%v]", c.Packages[0])
	}
	if c.Packages[1] != "p2" {
		t.Errorf("package load error p2 != [%v]", c.Packages[1])
	}
	if len(c.CommonFlags) != 3 {
		t.Errorf("CommonFlags size error 2 != [%v]", len(c.CommonFlags))
	}
	if c.CommonFlags[0] != "cf1" {
		t.Errorf("CommonFlags[0] error cf1 != [%v]", c.CommonFlags[0])
	}
	if _, ok := c.Entries["p1"]; !ok {
		t.Errorf("Package[p1] load error")
	}
	if _, ok := c.Entries["p2"]; !ok {
		t.Errorf("Package[p2] load error")
	}
	pkg1, _ := c.Entries["p1"]
	if pkg1.Location != "loc_p1" {
		t.Errorf("Package[p1] locatio load error loc_p1 != [%v]", pkg1.Location)
	}
	if len(pkg1.Patterns) != 2 {
		t.Errorf("Package[p1].Patterns size error 2 != [%v]", len(pkg1.Patterns))
	}
	if pkg1.Patterns[0] != "pat_p1_1" {
		t.Errorf("Package[p1].Patterns[0] load error pat_p1_1 != [%v]", pkg1.Patterns[0])
	}
	if pkg1.Targets[0] != "t_p1_1" {
		t.Errorf("Package[p1].Targets[0] load error t_p1_1 != [%v]", pkg1.Targets[0])
	}
	if pkg1.Flags[0] != "f_p1_1" {
		t.Errorf("Package[p1].Flags[0] load error f_p1_1 != [%v]", pkg1.Flags[0])
	}

	pkg2, _ := c.Entries["p2"]
	if pkg2.Location != "loc_p2" {
		t.Errorf("Package[p2] locatio load error loc_p2 != [%v]", pkg2.Location)
	}
	if pkg2.Patterns[0] != "pat_p2_1" {
		t.Errorf("Package[p2].Patterns[0] load error pat_p2_1 != [%v]", pkg2.Patterns[0])
	}
	if pkg2.Targets[0] != "t_p2_1" {
		t.Errorf("Package[p2].Targets[0] load error t_p2_1 != [%v]", pkg2.Targets[0])
	}
	if pkg2.Flags[0] != "f_p2_1" {
		t.Errorf("Package[p2].Flags[0] load error f_p2_1 != [%v]", pkg2.Flags[0])
	}
}
