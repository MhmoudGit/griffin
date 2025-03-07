package generators

import (
	"fmt"
	"griffin/internal/log"
	"griffin/internal/templates"
	"html/template"
	"os"
	"os/exec"
)

func NewProject(name string) error {
	log.Info("creating new project: ", name)

	if err := os.Mkdir(name, os.ModePerm); err != nil {
		return err
	}

	err := mainTemplate(name)
	if err != nil {
		return err
	}

	err = configTemplate(name)
	if err != nil {
		return err
	}

	err = gitignoreTemplate(name)
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

func mainTemplate(name string) error {
	mainFile, err := os.Create(name + "/main.go")
	if err != nil {
		return err
	}
	defer mainFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/main.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(mainFile, nil)
	if err != nil {
		return err
	}

	return nil
}

type ConfigData struct {
	Name string
}

func configTemplate(name string) error {
	configFile, err := os.Create(name + "/config.yaml")
	if err != nil {
		return err
	}
	defer configFile.Close()

	tmpl, err := template.ParseFS(templates.FS, "project/config.yaml.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	data := ConfigData{
		Name: name,
	}

	err = tmpl.Execute(configFile, data)
	if err != nil {
		return err
	}

	return nil
}

func gitignoreTemplate(name string) error {
	gitignoreFile, err := os.Create(name + "/.gitignore")
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
	if err := os.Chdir(projectDir); err != nil {
		return fmt.Errorf("failed to change to project directory: %v", err)
	}

	cmd := exec.Command("go", "mod", "init", projectDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize Go module: %v", err)
	}

	log.Success("Go module initialized successfully.")
	return nil
}

func installDependencies() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install dependencies: %v", err)
	}

	fmt.Println("Dependencies installed successfully.")
	return nil
}
