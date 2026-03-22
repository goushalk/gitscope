package config

import (
	// "errors"
	"errors"
	"os"
	"path/filepath"
)

const (
	AppName  = "gitscope"
	fileName = "config.json"
)

func DataDir() (string, error) { // defines the directory where the congig DIR should be.

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", AppName), nil

}

func ConfigFilePath() (string, error) { // defines the config file path.
	dataDir, err := DataDir()

	if err != nil {
		return "", err
	}

	return filepath.Join(dataDir, fileName), nil
}

func EnsureStorage() error { // Ensures that all the above directories and files exist.

	dataDir, err := DataDir()

	if err != nil {
		return errors.New("config folder not found, creating config folder")
	}
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return err
	}

	return nil
}

// TODO: Make the function to.

func CreateAndWriteFile(path string, data []byte) error {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
