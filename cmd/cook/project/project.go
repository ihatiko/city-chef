package project

import (
	"embed"
	"errors"
	dynamic_struct "github.com/ihatiko/dynamic-struct"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	templateKey = "template"
)

//go:embed template
var structure embed.FS

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "project",
		Short: "Generate project template by clean architecture",
		Run: func(cmd *cobra.Command, args []string) {
			folder := "C:\\Users\\user\\GolandProjects\\awesomeProject8"
			obj := dynamic_struct.ConstructStruct(map[string]any{})
			process(templateKey, folder, obj)
		},
	}
}

func process(prefix, destination string, obj any) {
	dir, err := structure.ReadDir(prefix)
	if err != nil {
		log.Fatal(err)
	}
	for _, fl := range dir {
		secondPath := filepath.Join(prefix, fl.Name())

		if fl.IsDir() {
			directory := cleanPath(secondPath)
			newDirectoryPath := filepath.Join(destination, directory)
			Mkdir(newDirectoryPath)

			if checkWindows() {
				secondPath = filepath.ToSlash(secondPath)
			}
			process(secondPath, destination, obj)
			continue
		}

		buildFile(secondPath, destination, obj)
	}
}
func checkWindows() bool {
	return os.Getenv("GOOS") == "windows" ||
		strings.Contains(strings.ToLower(os.Getenv("OS")), "windows")
}

func cleanPath(path string) string {
	if checkWindows() {
		path = filepath.ToSlash(path)
	}
	result, _ := strings.CutPrefix(path, "template/")
	return result
}

func buildFile(secondPath string, folder string, obj any) {
	if checkWindows() {
		secondPath = filepath.ToSlash(secondPath)
	}
	b, err := structure.ReadFile(secondPath)
	if err != nil {
		log.Fatal(err)
	}
	t, err := template.New("").Parse(string(b))
	if err != nil {
		log.Fatal(err)
	}
	parsedPath, _ := strings.CutSuffix(secondPath, ".tmpl")
	parsedPath = cleanPath(parsedPath)

	filePath := filepath.Join(folder, parsedPath)
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(f, "", obj)
	if err != nil {
		log.Fatal(err)
	}
}

func Mkdir(path string) {
	var err error
	_, err = os.ReadDir(path)
	if errors.Is(err, fs.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
