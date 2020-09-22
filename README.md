# Yaml Templater   
For easy use of YAML in bash environments:

It is a very small Go binary utility that just wraps ```text/template``` library of Golang to be able to format multiple commands, one for each entry of dockerImages.

<b>Why?</b>
- In some repositories we want to build multiple docker containers.
- Kaniko only has a busybox shell (no NodeJS, no Golang, nothing), so we need a bash script or a separate binary!
- In bash there is no easy way to parse ```.hue/deploy.yml#dockerImages```.

## Usage

How to use:
```bash
$ yaml_templater <file.yaml> <json-path-selector> <command including go expressions>
```

## Example

```bash
$ yaml_templater deploy.yml  "{{ range .dockerImages }} 
  echo {{ .image }} ; 
  echo {{ .context }} ;
  {{ end }}" | sh
```
