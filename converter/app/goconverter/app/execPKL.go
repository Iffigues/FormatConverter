package main

import (
	"os/exec"
	"fmt"
)

func (f *Format) ExecPKL(a string) error {
	cmd := exec.Command("/home/debian/pkl", "eval", "-f", a, f.PKLPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, string(output), "a = ", a , "f = ",f.PKLPath)
		return err
	}
	f.RenderPKL = output
	return nil
}
