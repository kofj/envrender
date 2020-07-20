package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var input, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.WriteString("read error: " + err.Error())
		return
	}

	envMap, _ := envToMap()
	t := template.Must(template.New("tmpl").Parse(string(input)))
	err = t.Execute(os.Stdout, envMap)
	if err != nil {
		os.Stderr.WriteString("render failed: " + err.Error())
	}
}

func envToMap() (map[string]string, error) {
	envMap := make(map[string]string)
	var err error

	for _, v := range os.Environ() {
		split := strings.Split(v, "=")
		envMap[split[0]] = split[1]
	}

	return envMap, err
}
