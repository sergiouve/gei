package installer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func RefreshLocalDatabase() {
	fmt.Println("refreshin database...")
}

func databaseFileIsSane() bool {
	return localDatabaseFileExists()
}

func localDatabaseFileExists() bool {
	homeDir, _ := os.UserHomeDir()
	fileName := "localDatabase.json"

	_, err := os.Open(filepath.Join(fmt.Sprintf("%s/.gei", homeDir), filepath.Base(fileName)))

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
