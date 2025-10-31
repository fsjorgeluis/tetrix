package infrastructure

import (
	"log"
	"os"
	"path/filepath"
)

func assetPath(rel string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get working directory: %v", err)
	}
	full := filepath.Join(wd, rel)
	return full
}
