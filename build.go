package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	cKey   = regexp.MustCompile("(?i)" + ".*\\.c$")
	cppKey = regexp.MustCompile("(?i)" + ".*\\.cpp$")
)

func buildMatchPatterns(c *LspConfig) []*regexp.Regexp {
	ret := []*regexp.Regexp{}
	for _, p := range c.Patterns {
		k := regexp.MustCompile(fmt.Sprintf("(?i)"+".*\\.%v$", p))
		ret = append(ret, k)
	}
	return ret
}

func findFiles(conf *LspConfig) ([]string, error) {
	keys := buildMatchPatterns(conf)
	ret := []string{}

	cur, _ := os.Getwd()
	defer os.Chdir(cur)

	for _, p := range conf.Targets {
		workPath := filepath.Join(conf.Location, p)
		if err := os.Chdir(workPath); err != nil {
			return nil, err
		}
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			for _, k := range keys {
				if k.MatchString(path) {
					//ret = append(ret, filepath.Join(workPath, path))
					ret = append(ret, filepath.Join(p, path))
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func buildStructure(conf *LspConfig, commonFlags []string) (*[]ClangdOptionEntry, error) {
	files, err := findFiles(conf)
	if err != nil {
		return nil, err
	}

	loc := filepath.ToSlash(conf.Location)
	f := conf.Flags
	f = append(f, commonFlags...)

	d := []ClangdOptionEntry{}
	for _, e := range files {
		d = append(d,
			ClangdOptionEntry{
				Directory: loc,
				Arguments: f,
				File:      filepath.ToSlash(e),
			})
	}
	return &d, nil

}
