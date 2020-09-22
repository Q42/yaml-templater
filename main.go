package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	//get command line arguments
	if len(os.Args) < 3 {
		log.Fatalf("Please provide at least 2 arguments:\n\n$ yaml-templater [file.yaml] [template expression]")
	}
	yamlFile := os.Args[1]
	jsonPath := os.Args[2]

	//load yaml file
	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	runTemplate(data, jsonPath, os.Stdout)
}

func runTemplate(data []byte, jsonPath string, dest io.Writer) {
	tmpl, err := template.New("temp").Parse(jsonPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//bind YAML data with unmarshal
	var object interface{}
	err = yaml.Unmarshal(data, &object)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	tmpl.Execute(dest, object)
}
