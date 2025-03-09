package generators

import (
	"fmt"
	"griffin/internal/log"
	"html/template"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateModel(resource string, fields []string) error {
	// Define the model template
	modelTemplate := `package {{.Package}}

import (
	"time"

	"gorm.io/gorm"
)

type {{.Resource}} struct {
	ID int ` + "`json:\"id\"`" + `
	{{- range .Fields }}
	{{.Name}} {{.Type}} ` + "`json:\"{{.Json}}\"`" + `
	{{- end }}
	CreatedAt    time.Time      ` + "`json:\"createdAt\"`" + `
	UpdatedAt    time.Time      ` + "`json:\"updateAt\"`" + `
	DeletedAt    gorm.DeletedAt ` + "`json:\"deletedAt\"`" + `
}
`

	// Map user-provided types to Go types
	typeMapping := map[string]string{
		"string": "string",
		"text":   "string",
		"int":    "int",
		"bool":   "bool",
		"float":  "float64",
		"time":   "time.Time",
	}

	// Parse the fields into a usable format
	type Field struct {
		Name string
		Type string
		Json string
	}
	var parsedFields []Field
	hasTime := false // Track if the model requires the "time" package

	for _, f := range fields {
		parts := strings.Split(f, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid field format: %s (expected 'name:type')", f)
		}

		fieldName := parts[0]
		fieldType := parts[1]

		// Map the user-provided type to a Go type
		goType, ok := typeMapping[fieldType]
		if !ok {
			return fmt.Errorf("unsupported field type: %s", fieldType)
		}

		// Check if the field requires the "time" package
		if goType == "time.Time" {
			hasTime = true
		}

		parsedFields = append(parsedFields, Field{
			Name: cases.Title(language.Und, cases.NoLower).String(fieldName), // Capitalize the field name
			Type: goType,
			Json: fieldName,
		})
	}

	// Create the model file
	modelFilePath := "model.go"
	modelFile, err := os.Create(modelFilePath)
	if err != nil {
		return fmt.Errorf("failed to create model file: %v", err)
	}
	defer modelFile.Close()

	// Execute the template
	tmpl, err := template.New("model").Parse(modelTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse model template: %v", err)
	}

	data := map[string]interface{}{
		"Package":  resource,
		"Resource": cases.Title(language.Und, cases.NoLower).String(resource), // Capitalize the resource name
		"Fields":   parsedFields,
		"HasTime":  hasTime,
	}

	if err := tmpl.Execute(modelFile, data); err != nil {
		return fmt.Errorf("failed to execute model template: %v", err)
	}

	log.Info("Generated model: ", modelFilePath)
	return nil
}
