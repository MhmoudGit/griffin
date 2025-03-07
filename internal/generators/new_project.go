package generators

import (
	"fmt"
	"griffin/internal/log"
	"griffin/internal/templates"
	"html/template"
	"os"
	"os/exec"
)

type Data struct {
	ProjectName string
}

func NewProject(name string) error {
	log.Info("creating new project: ", name)

	data := Data{
		ProjectName: name,
	}

	if err := os.Mkdir(name, os.ModePerm); err != nil {
		return err
	}

	if err := os.Chdir(name); err != nil {
		return fmt.Errorf("failed to change to project directory: %v", err)
	}

	err := mainTemplate(data)
	if err != nil {
		return err
	}

	err = databaseTemplate(data)
	if err != nil {
		return err
	}

	err = configTemplate(data)
	if err != nil {
		return err
	}

	err = jobsTemplate()
	if err != nil {
		return err
	}

	err = errorsTemplate()
	if err != nil {
		return err
	}

	err = gitignoreTemplate()
	if err != nil {
		return err
	}

	err = initGoModule(name)
	if err != nil {
		return err
	}

	err = installDependencies()
	if err != nil {
		return err
	}

	return nil
}

func mainTemplate(data Data) error {
	log.Info("creating /main.go")
	mainFile, err := os.Create("./main.go")
	if err != nil {
		return err
	}
	defer mainFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/main.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(mainFile, data)
	if err != nil {
		return err
	}

	return nil
}

func databaseTemplate(data Data) error {
	log.Info("creating /database/")
	if err := os.Mkdir("database", os.ModePerm); err != nil {
		return err
	}

	log.Info("creating /database/postgres.go")
	databaseFile, err := os.Create("./database/postgres.go")
	if err != nil {
		return err
	}
	defer databaseFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/database/postgres.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(databaseFile, data)
	if err != nil {
		return err
	}

	return nil
}

func configTemplate(data Data) error {
	log.Info("creating /config.yaml")
	configFile, err := os.Create("./config.yaml")
	if err != nil {
		return err
	}
	defer configFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/config.yaml.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(configFile, data)
	if err != nil {
		return err
	}

	log.Info("creating /config/")
	if err := os.Mkdir("config", os.ModePerm); err != nil {
		return err
	}

	err = configGoTemplate()
	if err != nil {
		return err
	}

	err = corsGoTemplate()
	if err != nil {
		return err
	}

	err = loggerGoTemplate()
	if err != nil {
		return err
	}

	err = validatorGoTemplate()
	if err != nil {
		return err
	}

	return nil
}

func configGoTemplate() error {
	log.Info("creating /config/config.go")

	configGoFile, err := os.Create("config/config.go")
	if err != nil {
		return err
	}
	defer configGoFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/config/config.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(configGoFile, nil)
	if err != nil {
		return err
	}

	return nil
}

func corsGoTemplate() error {
	log.Info("creating /config/cors.go")
	corsGoFile, err := os.Create("config/cors.go")
	if err != nil {
		return err
	}
	defer corsGoFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/config/cors.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(corsGoFile, nil)
	if err != nil {
		return err
	}

	return nil
}

func loggerGoTemplate() error {
	log.Info("creating /config/logger.go")
	loggerGoFile, err := os.Create("config/logger.go")
	if err != nil {
		return err
	}
	defer loggerGoFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/config/logger.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(loggerGoFile, nil)
	if err != nil {
		return err
	}

	return nil
}

func validatorGoTemplate() error {
	log.Info("creating /config/validator.go")
	validatorGoFile, err := os.Create("config/validator.go")
	if err != nil {
		return err
	}
	defer validatorGoFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/config/validator.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(validatorGoFile, nil)
	if err != nil {
		return err
	}

	return nil
}

func jobsTemplate() error {
	log.Info("creating /jobs/")
	if err := os.Mkdir("jobs", os.ModePerm); err != nil {
		return err
	}

	err := emailsGoTemplate()
	if err != nil {
		return err
	}

	err = uploadsGoTemplate()
	if err != nil {
		return err
	}

	return nil
}

func emailsGoTemplate() error {
	log.Info("creating /jobs/emails.go")
	emailsGoFile, err := os.Create("jobs/emails.go")
	if err != nil {
		return err
	}
	defer emailsGoFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/jobs/emails.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(emailsGoFile, nil)
	if err != nil {
		return err
	}

	return nil
}

func uploadsGoTemplate() error {
	log.Info("creating /jobs/uploads.go")
	uploadsGoFile, err := os.Create("jobs/uploads.go")
	if err != nil {
		return err
	}
	defer uploadsGoFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/jobs/uploads.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(uploadsGoFile, nil)
	if err != nil {
		return err
	}

	return nil
}

func errorsTemplate() error {
	log.Info("creating /errors/")
	if err := os.Mkdir("errors", os.ModePerm); err != nil {
		return err
	}

	log.Info("creating /errors/errors.go")
	errorsFile, err := os.Create("./errors/errors.go")
	if err != nil {
		return err
	}
	defer errorsFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/errors/errors.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(errorsFile, nil)
	if err != nil {
		return err
	}

	return nil
}

func gitignoreTemplate() error {
	log.Info("creating /.gitignore")
	gitignoreFile, err := os.Create("./.gitignore")
	if err != nil {
		return err
	}
	defer gitignoreFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/.gitignore.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(gitignoreFile, nil)
	if err != nil {
		return err
	}

	return nil
}

func initGoModule(projectDir string) error {
	cmd := exec.Command("go", "mod", "init", projectDir)
	cmd.Stdout = log.StdWriter
	cmd.Stderr = log.StdWriter

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize Go module: %v", err)
	}

	log.Success("Go module initialized successfully.")
	return nil
}

func installDependencies() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = log.StdWriter
	cmd.Stderr = log.StdWriter

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install dependencies: %v", err)
	}

	log.Success("dependencies installed successfully.")
	return nil
}
