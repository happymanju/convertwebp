package convertwebp

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var targetFileExtensions []string = []string{"jpg", "jpeg", "png"}

func Run(args []string) int {
	currentDir, _ := os.Getwd()
	err := os.Mkdir("webp", 0755)
	if err != nil {
		log.Printf("error making webp directory: %v\n", err)
		return 1
	}

	allFiles, err := os.ReadDir(currentDir)
	if err != nil && len(allFiles) == 0 {
		log.Printf("Error reading directory files: %v\n", err)
		return 1
	}

	fmt.Printf("Converting images in %s...\n", currentDir)

	for _, file := range allFiles {
		if !file.IsDir() {
			convertFlag := false
			var extIndex int
			for k, fileExtension := range targetFileExtensions {
				if strings.Contains(file.Name(), fileExtension) {
					convertFlag = true
					extIndex = k
				}
			}
			if convertFlag {
				sp := exec.Command("cwebp", "-progress", "-pass", "10", "-m", "6", file.Name(), "-o", filepath.Join("webp", strings.Replace(file.Name(), targetFileExtensions[extIndex], "webp", 1)))

				fmt.Printf("Converting: %s...\n", file.Name())
				err = sp.Run()
				if err != nil {
					log.Printf("error processing file %s: %v\n", file.Name(), err)
				}
			}
		}
	}

	return 0
}
