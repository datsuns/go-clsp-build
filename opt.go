package main

import "flag"

type Option struct {
	ConfPath string
	DestPath string
}

func parseOption() *Option {
	ret := &Option{}
	flag.StringVar(&ret.ConfPath, "c", "example.yml", "path to config file")
	flag.StringVar(&ret.DestPath, "d", "compile_commands.json", "path to dest file")
	flag.Parse()
	return ret
}
