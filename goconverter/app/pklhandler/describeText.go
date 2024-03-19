package pklhandler

import (
	"converter/data"
	"fmt"
	"net/http"
)

func (p *PklHandler) DescribeText(w http.ResponseWriter, r *http.Request) {

	e := `{
  "name": "John Doe",
  "age": 30,
  "is_student": false,
  "address": {
    "street": "123 Main Street",
    "city": "Anytown",
    "zip_code": "12345"
  },
  "email_addresses": [
    "john.doe@example.com",
    "johndoe@gmail.com"
  ]
}`
	fmt.Println(data.MagikaDescribeText(e))
}
