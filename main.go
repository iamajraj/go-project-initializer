package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"
)

var directoryPath string

func initialize() {
	if _, err := os.Stat("config.txt"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("config not found creating one...")
		file, err := os.Create("config.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println("config file has been created ✅")
		var directory string
		fmt.Print("Enter the default directory name: ")
		fmt.Scanln(&directory)

		if s := strings.Trim(directory, ""); s == "" {
			fmt.Println("Path not provided exiting...")
			os.Exit(1)
		}

		file.Write([]byte(fmt.Sprintf("path=%s", directory)))
		fmt.Println("config directory has been set!🤘")
	}

	data, err := os.ReadFile("config.txt")

	if err != nil {
		panic(err)
	}

	splitted := strings.Split(string(data), "=")

	directoryPath = strings.Trim(splitted[1], " ")
}

func main() {
	initialize()

	var projectName string

	fmt.Print("What's the name of the project ? -> ")
	fmt.Scanln(&projectName)

	if _, err := os.Open(directoryPath); os.IsExist(err) {
		fmt.Println("The directory named", projectName, "already exist ❌")
		os.Exit(0)
	}
	fmt.Printf("initializing project with name %s ✍\n", projectName)

	if err := os.Chdir(directoryPath); err != nil {
		panic(err)
	}

	err := os.Mkdir(projectName, fs.FileMode(os.O_CREATE))

	if err := os.Chdir(fmt.Sprintf("%s\\%s", directoryPath, projectName)); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Directory has been created ✨")

	fmt.Println("")
	init := exec.Command("go", "mod", "init", fmt.Sprintf("iamajraj/%s", projectName))
	_, err = init.Output()
	if err != nil {
		panic(err)
	}
	fileCreate := exec.Command("cmd", "/C", "type", "NUL", ">", "main.go")
	_, err = fileCreate.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	fmt.Println("Project initialized with go mod init 🚀")
	fmt.Println("opening directory in vscode...")
	open := exec.Command("code", ".")
	open.Output()
}
