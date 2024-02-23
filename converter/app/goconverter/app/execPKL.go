package main

import (
	"os/exec"
)

func (f *Format) ExecPKL(a string) error {
	cmd := exec.Command("/home/debian/pkl", "eval", "-f", a, f.PKLPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	f.RenderPKL = output
	return nil
}
