package main

import (
	"os/exec"
)


func (f *Format)ExecPKL(a string) ([]byte, error){
	cmd := exec.Command("/home/debian/pkl", "eval", "-f", a, f.PKLPath)
	output, err := cmd.CombinedOutput()
	return output, err
}
