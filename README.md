# Yaml Templater   

## Usage

How to use:
```bash
$ yaml_templater <file.yaml> <json-path-selector> -- <command including go expressions>
```

## Example

```bash
$ yaml_templater deploy.yml  "{{ range .dockerImages }} 
  echo {{ .image }} ; 
  echo {{ .context }} ;
  {{ end }}" | sh
```
