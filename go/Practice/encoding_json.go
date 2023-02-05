package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"course-price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	encodingJson()
}

func encodingJson() {
	const myurl string = ""

	courses := []course{
		{"C", 500, "surajmyt.code", "c@123", []string{"coding", "IT"}},
		{"Java", 800, "surajmyt.org", "java@123", []string{"dev", "IT"}},
		{"Kotlin", 1500, "surajmyt.dev", "kt@123", nil},
	}

	// packing data into json
	fmt.Println("Using json.Marshal():")
	json_pack, errJsonM := json.Marshal(courses)
	printErr(errJsonM)

	fmt.Println(string(json_pack))

	fmt.Printf("\n\n\n\n")

	fmt.Println("Using json.MarshalIndent():")

	json_pack2, errJsonM2 := json.MarshalIndent(courses, "", "\t")
	printErr(errJsonM2)

	fmt.Println(string(json_pack2))
}
func printErr(err error) {
	if err != nil {
		panic(err)
	}
}
