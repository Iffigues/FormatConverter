package main

import (
	"os/exec"
)

func (f *Format) ExecPKL(a, out  string) error {
	cmd := exec.Command("pkl", "eval", "-f", a, f.PKLPath, "-o", out)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
