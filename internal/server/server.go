package server

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/MhmoudGit/griffin/internal/log"
)

func Start() error {
	// Get the current working directory (project root)
	projectDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting project directory: %v", err)
	}

	// Change to the project directory
	err = os.Chdir(projectDir)
	if err != nil {
		return fmt.Errorf("failed to change to project directory: %v", err)
	}

	// Create the command to start the server
	cmd := exec.Command("go", "run", "main.go")

	// Capture the command's output and errors
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	log.Info("Starting server...")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	// Create a channel to listen for termination signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a termination signal or for the server to exit
	select {
	case <-sigChan:
		// Gracefully shut down the server
		log.Info("Received termination signal. Shutting down server...")

		// Handle graceful shutdown based on the operating system
		if runtime.GOOS == "windows" {
			if err := cmd.Process.Kill(); err != nil {
				return fmt.Errorf("failed to kill server process: %v", err)
			}
		} else {
			if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
				return fmt.Errorf("failed to send termination signal to server: %v", err)
			}
		}

		// Wait for the server to shut down
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		done := make(chan error, 1)
		go func() {
			done <- cmd.Wait()
		}()

		select {
		case err := <-done:
			if err != nil {
				return fmt.Errorf("server exited with error: %v", err)
			}
			fmt.Println("Server shut down gracefully.")
		case <-ctx.Done():
			return fmt.Errorf("server shutdown timed out: %v", ctx.Err())
		}
	case err := <-waitForCommand(cmd):
		if err != nil {
			return fmt.Errorf("server exited with error: %v", err)
		}
	}

	return nil
}

func waitForCommand(cmd *exec.Cmd) <-chan error {
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	return done
}
