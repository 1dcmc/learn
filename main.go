package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	Salary    int    `json:"salary"`
	fullTime  bool   `json:"fullTime"`
}

func main() {
	Ross := Employee{
		FirstName: "Ross",
		lastName:  "Green",
		Salary:    2000,
		fullTime:  true,
	}
	jsonBytes, err := json.Marshal(Ross)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Employee json:\n%s", string(jsonBytes))

}
