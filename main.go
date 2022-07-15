package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type jsonData struct {
	State string
}

func main() {

	var data []jsonData
	m := make(map[string]int)

	c := http.Client{}

	resp, err := c.Get("https://api.github.com/repos/google/material-design-icons/issues?state=all&per_page=100&page=1")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &data)

	for _, v := range data {
		if v.State == "open" {
			m["open"] += 1
		} else {
			m["closed"] += 1
		}
	}
	jsonmarshaled, _ := json.MarshalIndent(m, "", "  ")
	fmt.Print(string(jsonmarshaled))
}
