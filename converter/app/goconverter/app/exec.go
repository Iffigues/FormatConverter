package main

import (
	"fmt"
	"os/exec"
)

func (f *Format) Exec(a, out  string) error {
	dest := "/tmp/file/generatedpkl/" + f.pklDirname + "/" +a + ".pkl"
	cmd := exec.Command("pkl", "eval", dest, "-o", out)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err, la ", err, string(output), a)
		return err
	}
	f.PKLPath  =  dest
	return nil
}
