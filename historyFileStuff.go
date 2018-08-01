package historyFileStuff

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filePath, _ := filepath.Abs("history/history.txt")
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Something worked?")
	}
	file.Close()

	if _, err := os.Stat("history"); os.IsNotExist(err) {
		os.Mkdir("history", os.ModePerm)
	}

	historyFile, historyFileErr := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if historyFileErr != nil {
		fmt.Println(err.Error())
	}

	historyFile.WriteString("I have added some text\r\n")
	fmt.Println("End of it")
	defer historyFile.Close()
}
