package cred

import (
	"bufio"
	"errors"
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/defs"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

// getConfigDir Get directory for soring credentials
// On Windows: %APPDATA%
// On Linux: $HOME/.config
func getConfigDir() (string, error) {
	// Get folder for storing apps data
	var configFolder string
	if runtime.GOOS == "windows" {
		if configFolder = os.Getenv("APPDATA"); configFolder == "" {
			return "", errors.New("APPDATA env is not set")
		}
	} else {
		u, err := user.Current()
		if err != nil {
			return "", err
		}
		configFolder = filepath.Join(u.HomeDir, ".config")
	}

	// Return the path to the app folder
	return filepath.Join(configFolder, defs.AppShortName), nil
}

// getConfigFilePath Get path to the file where the credentials are stored
func getConfigFilePath() (string, error) {
	appFolder, err := getConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(appFolder, "creds.txt"), nil
}

// createFileIfNotExists Create the file for credentials if it does not exist
func createFileIfNotExists() error {
	credFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	// Check if the file exists
	if _, err := os.Stat(credFilePath); os.IsNotExist(err) {
		// Create the directory
		if err := os.MkdirAll(filepath.Dir(credFilePath), 0766); err != nil {
			return err
		}
		// Create the file
		if _, err := os.Create(credFilePath); err != nil {
			return err
		}
		// Set the permissions
		if err := os.Chmod(credFilePath, 0600); err != nil {
			return err
		}
	}
	return nil
}

// StoreSessionId Store the session id in the file
func StoreSessionId(id string) error {
	if err := createFileIfNotExists(); err != nil {
		return err
	}
	credFilePath, _ := getConfigFilePath()
	// Write the session id to the file
	file, err := os.OpenFile(credFilePath, os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(id); err != nil {
		return err
	}
	return nil
}

// DeleteSessionId Delete the session id from the file
func DeleteSessionId() error {
	credFilePath, _ := getConfigFilePath()
	if err := os.Remove(credFilePath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// GetSessionId Get the session id from the file
func GetSessionId() (string, error) {
	credFilePath, _ := getConfigFilePath()

	// Read the session id from the file
	file, err := os.OpenFile(credFilePath, os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Scan()
	return fs.Text(), nil
}
