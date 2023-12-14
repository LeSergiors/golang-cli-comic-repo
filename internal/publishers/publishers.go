package publishers

import (
	"comic-cli/internal/files"
	"encoding/json"
)

type publishers struct {
	Publishers []publisher `json:"publishers"`
}

type publisher struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Example() {
	s := files.ReadFileContent("publishers")

	var p publishers
	err := json.Unmarshal([]byte(s), &p)

	if err != nil {
		panic(err)
	}
}
