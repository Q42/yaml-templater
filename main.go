package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	//get command line arguments
	yamlFile := os.Args[1]
	jsonPath := os.Args[2]

	tmpl, err := template.New("temp").Parse(jsonPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//load yaml file
	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//bind YAML data with unmarshal
	var object interface{}
	err = yaml.Unmarshal(data, &object)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	tmpl.Execute(os.Stdout, object)
}
