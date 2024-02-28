package entx

import (
	"embed"
	"html/template"
	"log"
	"os"
	"strings"

	"entgo.io/ent/entc/gen"
	"github.com/stoewer/go-strcase"
)

var (
	//go:embed templates/*
	_templates embed.FS
)

// schema data for template
type schema struct {
	Name string
}

// GenSchema generates graphql schemas when not specified to be skipped
func GenSchema(graphSchemaDir string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			// create schema template
			tmpl := createTemplate()

			// loop through all nodes and generate schema if not specified to be skipped
			for _, node := range g.Nodes {
				// check skip annotation
				if sg, ok := node.Annotations[SchemaGenAnnotationName]; ok {
					val, _ := sg.(map[string]interface{})["Skip"]

					if val.(bool) {
						continue
					}
				}

				// check if schema already exists, skip generation so we don't overwrite manual changes
				if _, err := os.Stat(graphSchemaDir + strings.ToLower(node.Name) + ".graphql"); err == nil {
					continue
				}

				file, err := os.Create(graphSchemaDir + strings.ToLower(node.Name) + ".graphql")
				if err != nil {
					log.Fatalf("Unable to create file: %v", err)
				}

				s := schema{
					Name: node.Name,
				}

				// execute template and write to file
				if err = tmpl.Execute(file, s); err != nil {
					log.Fatalf("Unable to execute template: %v", err)
				}
			}

			return next.Generate(g)
		})
	}
}

// createTemplate creates a new template for generating graphql schemas
func createTemplate() *template.Template {
	// function map for template
	fm := template.FuncMap{
		"ToLowerCamel": strcase.LowerCamelCase,
	}

	// create schema template
	tmpl, err := template.New("graph.tpl").Funcs(fm).ParseFS(_templates, "templates/graph.tpl")
	if err != nil {
		log.Fatalf("Unable to parse template: %v", err)
	}

	return tmpl
}
