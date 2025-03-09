package generators

import (
	"fmt"
	"griffin/internal/log"
	"html/template"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateRepository(resource string) error {
	// Define the repository template
	repositoryTemplate := `package {{.Package}}

import (
	"gorm.io/gorm"
)

type {{.Resource}}Repository struct {
	db *gorm.DB
}

func New{{.Resource}}Repository(db *gorm.DB) *{{.Resource}}Repository {
	return &{{.Resource}}Repository{db: db}
}

// List retrieves all {{.Resource}}s from the database
func (r *{{.Resource}}Repository) List({{.Resource}}s *[]{{.Resource}}) error {
	return r.db.Find({{.Resource}}s).Error
}

// Create adds a new {{.Resource}} to the database
func (r *{{.Resource}}Repository) Create({{.Resource}} *{{.Resource}}) error {
	return r.db.Create({{.Resource}}).Error
}

// Get retrieves a {{.Resource}} by ID from the database
func (r *{{.Resource}}Repository) Get(id uint, {{.Resource}} *{{.Resource}}) error {
	return r.db.First({{.Resource}}, id).Error
}

// Update modifies an existing {{.Resource}} in the database
func (r *{{.Resource}}Repository) Update({{.Resource}} *{{.Resource}}) error {
	return r.db.Save({{.Resource}}).Error
}

// Delete removes a {{.Resource}} from the database by ID
func (r *{{.Resource}}Repository) Delete(id uint) error {
	return r.db.Delete(&{{.Resource}}{}, id).Error
}
`

	// Create the repository file
	repositoryFilePath := "repository.go"
	repositoryFile, err := os.Create(repositoryFilePath)
	if err != nil {
		return fmt.Errorf("failed to create repository file: %v", err)
	}
	defer repositoryFile.Close()

	// Execute the template
	tmpl, err := template.New("repository").Parse(repositoryTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse repository template: %v", err)
	}

	data := map[string]interface{}{
		"Package":  resource,
		"Resource": cases.Title(language.Und, cases.NoLower).String(resource), // Capitalize the resource name
	}

	if err := tmpl.Execute(repositoryFile, data); err != nil {
		return fmt.Errorf("failed to execute repository template: %v", err)
	}

	log.Info("Generated repository: ", repositoryFilePath)
	return nil
}
