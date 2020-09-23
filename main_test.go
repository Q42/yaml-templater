package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var deployYaml = strings.TrimSpace(`
name: app-kustomize
kustomize: . # directory relative to PWD!
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
	assert.Equal(t, output.String(), expectedOutput)
}
