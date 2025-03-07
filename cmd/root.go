/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "griffin",
	Short: "Griffin: A powerful CLI tool for scaffolding Go web applications with Echo and GORM",
	Long: `Griffin is a command-line tool designed to streamline the development of Go web applications. 
Inspired by the mythical Griffin—a creature of strength and agility—this tool empowers developers 
to quickly generate CRUD (Create, Read, Update, Delete) applications with minimal effort.

Griffin integrates seamlessly with the Echo web framework and GORM ORM, providing a robust 
foundation for building scalable and maintainable web applications. With Griffin, you can:

- Generate models, handlers, routes, and migrations in seconds.
- Automate database setup and configuration.
- Scaffold RESTful APIs with JSON responses.
- Follow best practices for Go web development.

Examples:
  # Generate a CRUD application for a "posts" resource
  griffin generate resource posts title:string body:text

  # Start the development server
  griffin start

  # Run database migrations
  griffin migrate

Griffin is the ultimate tool for Go developers who want to focus on building great applications 
without getting bogged down by boilerplate code. Let Griffin handle the heavy lifting while you 
soar to new heights in your development journey.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.griffin.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
