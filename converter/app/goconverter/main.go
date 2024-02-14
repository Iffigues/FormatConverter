package main

import (
	"fmt"
	"io/ioutil"
	"net/http"	
)

func Home(w http.ResponseWriter, r *http.Request) {

		// Read the content of the HTML file
		content, err := ioutil.ReadFile("./static/index.html")
		if err != nil {
			http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to "text/html"
		w.Header().Set("Content-Type", "text/html")

		// Write the HTML content to the response
		w.Write(content)

}

func main() {
	ee := NewFormat("json", "./tmp/file/file.json")
	fmt.Println(ee.CreateFile())
	fmt.Println(ee.Exec())
	ee.CreatePKL()
	j, _ := ee.ExecPKL("json")
	fmt.Println(ee.CreateNew(j))
	ee.Erase()
	return
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8780", nil)
}
