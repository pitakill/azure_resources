// +build ignore
package main

import (
	"os"
	"strings"
	"text/template"
	"time"
)

//go:generate go run gen.go helpers.go
//go:generate gofmt -s -w ../

var (
	resources = []string{
		"Microsoft.Compute/virtualMachines",
		"Microsoft.Compute/virtualMachines/Extensions",
		"Microsoft.Network/publicIPAddresses",
		"Microsoft.Compute/disks",
		"Microsoft.Network/virtualNetworks",
		// This resource does not compliant struct of the
		//"Microsoft.Resources/deployments",
		// This resource does not exist in the Go Azure SDK client spec
		//"Microsoft.Network/networkInterfaces",
	}
)

func main() {
	template := template.Must(template.New("").Funcs(tplFuncs).ParseGlob("*.tpl"))

	generateResources(template)
	generateMain(template)
}

func generateResources(t *template.Template) {
	rscs := make(map[string][]string)

	for _, r := range resources {
		typ, resource := getTypeAndResource(r)
		rscs[typ] = append(rscs[typ], resource)
	}

	for typ, r := range rscs {
		f, err := os.Create("../" + strings.ToLower(typ) + ".go")
		die(err)
		defer f.Close()

		err = t.ExecuteTemplate(f, "resource.tpl", struct {
			Resources []string
			Timestamp time.Time
			Type      string
		}{
			Resources: r,
			Timestamp: now,
			Type:      typ,
		})
		die(err)
	}
}

func generateMain(t *template.Template) {
	f, err := os.Create("../main.go")
	die(err)
	defer f.Close()

	err = t.ExecuteTemplate(f, "main.tpl", struct {
		Timestamp time.Time
		Resources []string
	}{
		Timestamp: now,
		Resources: resources,
	})
	die(err)
}
