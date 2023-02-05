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
	decodingJson()
}
func decodingJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "Kotlin BootCamp",
			"Price": 500,
			"Platform": "kt.dev",
			"tags": ["android","development"]
		}
	`)

	var courses course

	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		fmt.Println("Valid Json")
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("\n\n%#v\n\n", courses)
	} else {
		fmt.Println("Invalid Json")
	}


	// ADD DATA TO KEY VALUE
	var fetchDataMap map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &fetchDataMap)
	fmt.Printf("\n\n%#v\n\n", fetchDataMap)


	// LOOP THOUGH MAP
	for key, val := range fetchDataMap {
		fmt.Printf("[%v] : [%v]{type: %T}\n",key, val, val)
	}
}
