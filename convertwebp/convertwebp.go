package convertwebp

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var targetFileExtensions []string = []string{".jpg", ".jpeg", ".png"}

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

	var images map[string]string = make(map[string]string)

	for _, file := range allFiles {
		if !file.IsDir() {
			for _, v := range targetFileExtensions {
				if strings.Contains(file.Name(), v) {
					images[file.Name()] = v
				}
			}
		}
	}

	fmt.Printf("Converting images in %s...\n", currentDir)
	fmt.Printf("Found %d image files\n", len(images))

	counter := 1

	for file, ext := range images {
		sp := exec.Command("cwebp", "-progress", "-pass", "10", "-m", "6", file, "-o", filepath.Join("webp", strings.Replace(file, ext, ".webp", 1)))

		fmt.Printf("Converting %d of %d: %s...\n", counter, len(images), file)
		err = sp.Run()
		if err != nil {
			log.Printf("error processing file %s: %v\n", file, err)
		}
		counter++

	}

	return 0
}
