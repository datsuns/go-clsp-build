package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func logging(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func execute(bin string, params ...string) error {
	cmd := exec.Command(bin, params...)
	e, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return err
	}
	o, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return err
	}

	go logging(e)
	go logging(o)
	cmd.Start()
	return cmd.Wait()
}
