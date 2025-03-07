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

	err = configTemplate(data)
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

	err = configGoTemplate()
	if err != nil {
		return err
	}

	return nil
}

func configGoTemplate() error {
	log.Info("creating /config/config.go")
	if err := os.Mkdir("config", os.ModePerm); err != nil {
		return err
	}

	configGoFile, err := os.Create("config/config.go")
	if err != nil {
		return err
	}
	defer configGoFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/config.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(configGoFile, nil)
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
