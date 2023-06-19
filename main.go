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

	j2, err := jinja2.NewJinja2("example", 1,
		jinja2.WithGlobal(strings.Split(variables, "=")[0],
			strings.Split(variables, "=")[1])) // no vars yet
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
