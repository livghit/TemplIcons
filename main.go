package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// this generates Templ files from the SVG files found under icons/outline folder
func main() {
	err := filepath.Walk("icons/outline", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".svg" {
			// Open the SVG file
			svgFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer svgFile.Close()

			// Read the SVG file content
			svgContent, err := io.ReadAll(io.Reader(svgFile))
			if err != nil {
				return err
			}

			// Get the icon name and convert it to camel case without - and _
			iconName := strings.ReplaceAll(info.Name(), "-", " ")
			iconName = strings.Title(iconName)
			iconName = strings.ReplaceAll(iconName, " ", "")
			iconName = strings.ReplaceAll(iconName, ".Svg", "")

			// Create the Go template file
			goFile, err := os.Create(fmt.Sprintf("heroicons/%s.templ", iconName))
			if err != nil {
				return err
			}
			defer goFile.Close()

			// Write the SVG content into the Go template file
			fmt.Fprintf(goFile, "package heroicons\n\ntempl %s(){\n%s\n}\n", iconName, strings.ReplaceAll(string(svgContent), `"/>`, `"></path>`))
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", "outline", err)
	}
}
