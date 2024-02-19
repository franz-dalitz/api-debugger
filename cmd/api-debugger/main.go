package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/franz-dalitz/api-debugger/yaml"
)

func main() {
	config, err := yaml.LoadConfig("examples/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(bytes))
}
