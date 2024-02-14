package main

import (
	"fmt"
	"os/exec"
)


func (f *Format)Exec() (error)  {
	fmt.Println("./tmp/pkl/"+f.id + ".pkl")
	cmd := exec.Command("/home/debian/pkl", "eval", "./tmp/pkl/" + f.id + ".pkl")
	output, err := cmd.CombinedOutput()
	if err  != nil {
		return err
	}
	f.PKL = output
	return nil
} 
