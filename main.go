package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type person struct {
	Name  string
	Age   int
	Color string
}

var logFile *os.File

func init() {
	logPath, _ := filepath.Abs("logs/logs.txt")
	logFile, _ := os.OpenFile(logPath, os.O_APPEND, 0666)

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(logFile)

	log.Info("We created a log")
}

func main() {
	filePath, err := filepath.Abs("json-folder/file.json")
	if err != nil {
		log.Error("Error setting path,", err)
		return
	}
	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()

	if err != nil {
		log.Error("Error opening file,", err)
		return
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var person person
	json.Unmarshal(byteValue, &person)
	log.Info(person.Name)

	fmt.Println("hello world")

	defer logFile.Sync()
	defer logFile.Close()
}
