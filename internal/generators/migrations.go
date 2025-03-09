package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/MhmoudGit/griffin/internal/log"
)

func GenerateMigration(resource string, fields []string) error {
	// Define the migration template
	migrationTemplate := `-- +goose Up
CREATE TABLE {{.Resource}} (
    id SERIAL PRIMARY KEY,
    {{- range .Fields }}
    {{.Name}} {{.SQLType}} NOT NULL,
    {{- end }}
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE {{.Resource}};
`

	// Map user-provided types to SQL types
	sqlTypeMapping := map[string]string{
		"string": "TEXT",             // PostgreSQL: TEXT, MySQL: TEXT
		"text":   "TEXT",             // PostgreSQL: TEXT, MySQL: TEXT
		"int":    "INTEGER",          // PostgreSQL: INTEGER, MySQL: INT
		"bool":   "BOOLEAN",          // PostgreSQL: BOOLEAN, MySQL: BOOLEAN
		"float":  "DOUBLE PRECISION", // PostgreSQL: DOUBLE PRECISION, MySQL: DOUBLE
		"time":   "TIMESTAMP",        // PostgreSQL: TIMESTAMP, MySQL: TIMESTAMP
	}

	// Parse the fields into a usable format
	type Field struct {
		Name    string
		Type    string
		SQLType string
	}
	var parsedFields []Field
	for _, f := range fields {
		parts := strings.Split(f, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid field format: %s (expected 'name:type')", f)
		}

		// Map the Go type to an SQL type
		sqlType, ok := sqlTypeMapping[parts[1]]
		if !ok {
			return fmt.Errorf("unsupported SQL type for field: %s", parts[1])
		}

		parsedFields = append(parsedFields, Field{
			Name:    parts[0],
			SQLType: sqlType,
		})
	}

	// Create the migrations directory if it doesn't exist
	migrationsDir := "migrations"
	if err := os.MkdirAll(migrationsDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create migrations directory: %v", err)
	}

	// Generate the migration file name
	timestamp := time.Now().Format("20060102150405")
	migrationFileName := fmt.Sprintf("%s_create_%s_table.sql", timestamp, strings.ToLower(resource))
	migrationFilePath := filepath.Join(migrationsDir, migrationFileName)

	// Create the migration file
	migrationFile, err := os.Create(migrationFilePath)
	if err != nil {
		return fmt.Errorf("failed to create migration file: %v", err)
	}
	defer migrationFile.Close()

	// Execute the template
	tmpl, err := template.New("migration").Parse(migrationTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse migration template: %v", err)
	}

	data := map[string]interface{}{
		"Resource": strings.ToLower(resource),
		"Fields":   parsedFields,
	}

	if err := tmpl.Execute(migrationFile, data); err != nil {
		return fmt.Errorf("failed to execute migration template: %v", err)
	}

	log.Info("Generated migration file:", migrationFilePath)
	return nil
}
