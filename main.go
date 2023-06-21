package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"strings"

	"github.com/kluctl/go-jinja2"
)

var template string
var variables string
var content string

func init() {
	flag.StringVar(&template, "t", "", "Jinja2 Template or path to a template file")
	flag.StringVar(&variables, "v", "", "Variables for Jinja2 template.")
}

func processVarsIntoMap(vars string) map[string]any {
	var m map[string]any
	var ss []string

	ss = strings.Split(vars, ",")
	m = make(map[string]any)
	for _, pair := range ss {
		z := strings.Split(pair, "=")
		m[z[0]] = z[1]
	}
	return m
}

func main() {
	flag.Parse()

	if _, err := os.Stat(template); errors.Is(err, os.ErrNotExist) {
		content = template
	} else {
		temp, err := os.ReadFile(template)
		if err != nil {
			panic(err)
		}
		content = strings.TrimSpace(string(temp))
	}

	vars_map := processVarsIntoMap(variables)

	j2, err := jinja2.NewJinja2("example", 1,
		jinja2.WithGlobals(vars_map))
	if err != nil {
		panic(err)
	}
	defer j2.Close()

	s, err := j2.RenderString(content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("template: %s\nresult: %s", content, s)
}
