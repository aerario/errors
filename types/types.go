package main

import (
	"log"
	"os"
	"text/template"
)

//go:generate go run types.go

func main() {
	var types = struct {
		Types []string
	}{
		Types: []string{
			"Authentication",
			"Authorization",
			"BadRequest",
			"Validation",
			"NotFound",
			"AlreadyExists",
			"LimitExceeded",
			"Inconsistent",
			"Persistence",
			"Infrastructure",
			"ThirdParties",
			"Timeout",
		},
	}

	for templateName, codeFileName := range map[string]string{
		"./types.gentpl": "../kind.go",
		"./tests.gentpl": "../kind_test.go",
	} {
		var tmpl, err = template.ParseFiles(templateName)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(1)
		}

		var file *os.File
		if file, err = os.Create(codeFileName); err != nil {
			log.Printf("%v", err)
			os.Exit(1)
		}

		if err = tmpl.Execute(file, types); err != nil {
			log.Printf("%v", err)
			os.Exit(1)
		}
	}
}
