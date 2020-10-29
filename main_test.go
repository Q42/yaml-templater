package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var deployYaml = strings.TrimSpace(`
name: app
dockerImages:
- image: main
  context: .
- image: child
  context: subdir
`)

var dockerTemplate = "{{range .dockerImages}}docker build -t {{.image}}:latest {{ .context }};\n{{end}}"

func TestFixtureDockerImages(t *testing.T) {
	output := bytes.NewBuffer(nil)
	runTemplate([]byte(deployYaml), dockerTemplate, output)

	expectedOutput := strings.TrimPrefix(`
docker build -t main:latest .;
docker build -t child:latest subdir;
`, "\n")
	assert.Equal(t, expectedOutput, output.String())
}

var blogYaml = strings.TrimSpace(`
title: this-funny-thing
date: 2020-01-01 10:00:00
text: "Lorum ipsum dolor amet <strong>foobar</strong>!"

`)

func TestFixtureBlog(t *testing.T) {
	output := bytes.NewBuffer(nil)

	runTemplate([]byte(blogYaml), `<h1>{{html .title}}</h1><p>{{html .text}}</p>`, output)
	expectedOutput := strings.TrimPrefix(`<h1>this-funny-thing</h1><p>Lorum ipsum dolor amet &lt;strong&gt;foobar&lt;/strong&gt;!</p>`, "\n")
	assert.Equal(t, expectedOutput, output.String())
	output.Reset()

	runTemplate([]byte(blogYaml), `<h1>{{.title}}</h1><p>{{.text}}</p>`, output)
	expectedOutput = strings.TrimPrefix(`<h1>this-funny-thing</h1><p>Lorum ipsum dolor amet <strong>foobar</strong>!</p>`, "\n")
	assert.Equal(t, expectedOutput, output.String())
}
