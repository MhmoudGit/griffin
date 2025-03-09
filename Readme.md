# Griffin CLI
Griffin is a powerful command-line tool designed to streamline the development of Go web applications. Inspired by the mythical Griffin—a creature of strength and agility—this tool empowers developers to quickly generate CRUD (Create, Read, Update, Delete) applications with minimal effort. 

Griffin integrates seamlessly with the Echo web framework and GORM ORM, providing a robust foundation for building scalable and maintainable web applications.

## Features
- Project Scaffolding: Quickly bootstrap a new project with a single command.

- CRUD Generation: Generate models, handlers, repositories, and routes for any resource.

- Database Migrations: Automate database setup and configuration using Goose.

- RESTful APIs: Scaffold RESTful APIs with JSON responses.

- Best Practices: Follow best practices for Go web development.

## Installation
To install Griffin, ensure you have Go installed on your system, then run:
```bash
go install github.com/MhmoudGit/griffin/cmd/griffin@latest
```
## Usage
Griffin provides a set of commands to help you manage your Go web application:

## Create a New Project
```bash
griffin new my_project
```
This command will create a new project directory with the necessary structure to get started.

## Create a database
Make sure to create a postgres database with the same name as the new project's name, or change the config.yaml after project generation

## Generate CRUD Code for a Resource
```bash 
griffin gen posts title:string body:string
```
This command will generate CRUD code for a posts resource, including a model, handler, repository, and routes.

## Run Database Migrations
```bash
griffin migrate
```
This command will run database migrations using Goose.

## Start the Development Server
```
griffin server
```
This command will start the development server.

## Acknowledgments
- Inspired by the Elixir Phoenix CLI tools.
- Built with Cobra, Echo, and GORM.

<b>Griffin<b>: <q> The ultimate tool for Go developers who want to focus on building great applications without getting bogged down by boilerplate code. Let Griffin handle the heavy lifting while you soar to new heights in your development journey.
