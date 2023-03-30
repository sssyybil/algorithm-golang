package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	parentDir := "./bank/questions/"

	dirName := flag.String("dName", "000", "Please type question number.")
	qName := flag.String("qName", "bar", "Please type question name.")

	flag.Parse()

	err := os.Mkdir(parentDir+fmt.Sprintf("q%s", *dirName), 0755)
	if err != nil {
		log.Fatalln(err)
	}

	dirPath := parentDir + "q" + *dirName + "/"

	file, err := os.Create(dirPath + *qName + ".go")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data := fmt.Sprintf("package main\n /**\n * 【】\n * 🪢\n * @%s\n */", time.Now().Format("2006-01-02 15:04"))
	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln(err)
	}
}
