// +build ignore
package main

import (
	"io/ioutil"
	"log"
	"strings"
	"text/template"
	"time"
)

var (
	now      = time.Now()
	tplFuncs = template.FuncMap{
		"Title":            strings.Title,
		"ToLower":          strings.ToLower,
		"GetType":          getType,
		"GetResourceTitle": getResourceTitle,
		"GetParameters":    getParameters,
	}
)

func openFile(input string) []byte {
	f, err := ioutil.ReadFile(input)
	die(err)

	return f
}

func getType(s string) string {
	return strings.TrimLeft(strings.Split(s, "/")[0], "Microsoft.")
}

func getResource(s string) string {
	splitted := strings.Split(s, "/")
	if len(splitted) == 3 {
		return singularize(splitted[1]) + splitted[2]
	}

	return splitted[1]
}

func singularize(s string) string {
	return strings.TrimRight(s, "s")
}

func getResourceTitle(s string) string {
	return strings.Title(getResource(s))
}

func getTypeAndResource(s string) (string, string) {
	return getType(s), getResource(s)
}

func getParameters(s string) string {
	args := "context.Background(), r.groupName, *r.resourceName, \"\""

	switch s {
	case "disks":
		args = strings.TrimRight(args, ", \"\"")
	case "virtualMachineExtensions":
		args = args + ", \"\""
	}

	return args
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
