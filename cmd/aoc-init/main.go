package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalln("Please enter one argument for the day like: aoc-init 1")
	}

	dayNumber := args[1]
	folderName := fmt.Sprint("day", dayNumber)

	createFolder(folderName)

	mainTmp := parseTemplate("main-template.txt")
	mainTestTmp := parseTemplate("main_test-template.txt")

	mainFileName := fmt.Sprint("./", folderName, "/", "main.go")
	mainFile := createFile(mainFileName)
	createTempFile(mainTmp, mainFile, dayNumber)

	mainTestFileName := fmt.Sprint("./", folderName, "/", "main_test.go")
	mainTestFile := createFile(mainTestFileName)
	createTempFile(mainTestTmp, mainTestFile, dayNumber)

	inputTxtFileName := fmt.Sprint("./", folderName, "/", "input.txt")
	_ = createFile(inputTxtFileName)

	log.Println("created all files")
}

func parseTemplate(fileName string) *template.Template {
	tmp, err := template.ParseFiles(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	return tmp
}

func createTempFile(tmp *template.Template, file *os.File, data any) {
	err := tmp.Execute(file, data)
	if err != nil {
		log.Fatalln(err)
	}
}

func createFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	return file
}

func createFolder(name string) {
	err := os.Mkdir(name, os.FileMode(0755))
	if err != nil {
		log.Fatalln(err)
	}
}
