package main

import (
	"html/template"
	"io/ioutil"
)

func parseFile(filename string) (*template.Template, error) {
	f, err := Assets.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	t := template.New(filename)
	return t.Parse(string(b))
}
