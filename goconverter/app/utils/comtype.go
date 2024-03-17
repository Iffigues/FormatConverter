package utils

import "fmt"

func GetCompType(e []string) (r string) {
	fmt.Println(e)
	if len(e) == 4 {
		return e[3]
	}
	return "zip"
}
