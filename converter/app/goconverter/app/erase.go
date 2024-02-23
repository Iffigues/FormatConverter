package main

import "os"

func (f *Format) Erase() {
	os.Remove("./tmp/pkl/" + f.id + ".pkl")
	os.Remove(f.NewPath)
}
