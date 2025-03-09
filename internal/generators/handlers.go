package generators

import (
	"fmt"
	"griffin/internal/log"
	"html/template"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateHandlers(resource, projectName string) error {
	// Define the handler template
	handlerTemplate := `package {{.Package}}

import (
	"strconv"
	"net/http"

	"github.com/labstack/echo/v4"
	"{{.ProjectName}}/errors"
)

type {{.Resource}}Handler struct {
	repo *{{.Resource}}Repository
}

func New{{.Resource}}Handler(repo *{{.Resource}}Repository) *{{.Resource}}Handler {
	return &{{.Resource}}Handler{repo: repo}
}

// List retrieves all {{.Resource}}s
func (h *{{.Resource}}Handler) List(c echo.Context) error {
	var {{.Resource}}s []{{.Resource}}
	if err := h.repo.List(&{{.Resource}}s); err != nil {
		return errors.InternalServerErr(err.Error())
	}
	return c.JSON(http.StatusOK, {{.Resource}}s)
}

// Create adds a new {{.Resource}}
func (h *{{.Resource}}Handler) Create(c echo.Context) error {
	var {{.Resource}} {{.Resource}}
	if err := c.Bind(&{{.Resource}}); err != nil {
		return errors.BadRequestErr(err.Error())
	}
	if err := h.repo.Create(&{{.Resource}}); err != nil {
		return errors.InternalServerErr(err.Error())
	}
	return c.JSON(http.StatusCreated, {{.Resource}})
}

// Get retrieves a {{.Resource}} by ID
func (h *{{.Resource}}Handler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.BadRequestErr("invalid ID")
	}
	var {{.Resource}} {{.Resource}}
	if err := h.repo.Get(uint(id), &{{.Resource}}); err != nil {
		return errors.NotFoundErr("{{.Resource}} not found")
	}
	return c.JSON(http.StatusOK, {{.Resource}})
}

// Update modifies an existing {{.Resource}}
func (h *{{.Resource}}Handler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.BadRequestErr("invalid ID")
	}
	var {{.Resource}} {{.Resource}}
	if err := h.repo.Get(uint(id), &{{.Resource}}); err != nil {
		return errors.NotFoundErr("{{.Resource}} not found")
	}
	if err := c.Bind(&{{.Resource}}); err != nil {
		return errors.BadRequestErr(err.Error())
	}
	if err := h.repo.Update(&{{.Resource}}); err != nil {
		return errors.InternalServerErr(err.Error())
	}
	return c.JSON(http.StatusOK, {{.Resource}})
}

// Delete removes a {{.Resource}} by ID
func (h *{{.Resource}}Handler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.BadRequestErr("invalid ID")
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		return errors.InternalServerErr(err.Error())
	}
	return c.JSON(http.StatusOK, "{{.Resource}} deleted")
}
`

	// Create the handlers file
	handlerFilePath := "handlers.go"
	handlerFile, err := os.Create(handlerFilePath)
	if err != nil {
		return fmt.Errorf("failed to create handler file: %v", err)
	}
	defer handlerFile.Close()

	// Execute the template
	tmpl, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse handler template: %v", err)
	}

	data := map[string]interface{}{
		"Package": resource,
		"Resource":    cases.Title(language.Und, cases.NoLower).String(resource), // Capitalize the resource name
		"ProjectName": projectName,
	}

	if err := tmpl.Execute(handlerFile, data); err != nil {
		return fmt.Errorf("failed to execute handler template: %v", err)
	}

	log.Info("Generated handler: ", handlerFilePath)
	return nil
}
