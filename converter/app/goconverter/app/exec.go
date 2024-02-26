package main

import (
	"fmt"
	"os/exec"
)

func (f *Format) Exec(a string) error {
	cmd := exec.Command("/home/debian/pkl", "eval", "/tmp/file/generatedpkl/"+ a +".pkl")
	output, err := cmd.CombinedOutput()
	fmt.Println("f = llll ",f.PKLPath, output)
	if err != nil {
		fmt.Println("err, la ", err, string(output), a)
		return err
	}
	f.PKLPath  =  "/tmp/file/generatedpkl/"+ a +".pkl"  
	 fmt.Println("hazi ", f.PKLPath)
	f.PKL = output
	return nil
}
