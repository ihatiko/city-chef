package main

import (
	"city-chef/cmd"
	"os"
)

func main() {
	os.Args = append(os.Args, "cook", "project")
	cmd.Execute()
}
