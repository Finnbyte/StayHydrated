package utils

import (
	"path/filepath"
	"os"
)

func GetStayHydratedPath() string {
	path, err := filepath.Abs("./StayHydrated.exe")
	if err != nil {
		os.Exit(1)
	}
	
	return path
}