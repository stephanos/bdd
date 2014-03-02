package bdd

import (
	"101loops/json"
	"io/ioutil"
)

func LoadJSON(path string) *json.Json {
	bytes, err := ioutil.ReadFile("fixtures/" + path + ".json")
	if err != nil {
		panic(err)
	}

	js, err := json.New(bytes)
	if err != nil {
		panic(err)
	}

	return js
}
