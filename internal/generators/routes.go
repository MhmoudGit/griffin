package generators

import (
	"fmt"
	"griffin/internal/log"
	"html/template"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateRoutes(resource string) error {
	// Define the routes template
	routesTemplate := `package {{.Package}}

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetupRoutes sets up the CRUD routes for a {{.Package}}
func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// Initialize the repository and handler
	{{.Resource}}Repo := New{{.Resource}}Repository(db)
	{{.Resource}}Handler := New{{.Resource}}Handler({{.Resource}}Repo)

	// Register the routes
	e.GET("/{{.Package}}", {{.Resource}}Handler.List)
	e.POST("/{{.Package}}", {{.Resource}}Handler.Create)
	e.GET("/{{.Package}}/:id", {{.Resource}}Handler.Get)
	e.PUT("/{{.Package}}/:id", {{.Resource}}Handler.Update)
	e.DELETE("/{{.Package}}/:id", {{.Resource}}Handler.Delete)
}
	`

	// Create the handlers file
	routesFilePath := "routes.go"
	routesFile, err := os.Create(routesFilePath)
	if err != nil {
		return fmt.Errorf("failed to create handler file: %v", err)
	}
	defer routesFile.Close()

	// Execute the template
	tmpl, err := template.New("routes").Parse(routesTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse routes template: %v", err)
	}

	data := map[string]interface{}{
		"Package":  resource,
		"Resource": cases.Title(language.Und, cases.NoLower).String(resource), // Capitalize the resource name
	}

	if err := tmpl.Execute(routesFile, data); err != nil {
		return fmt.Errorf("failed to execute routes template: %v", err)
	}

	log.Info("Generated routes: ", routesFilePath)
	return nil
}
