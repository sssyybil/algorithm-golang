package utils

import (
	"algorithm-golang/types"
	"fmt"
	"log"
	"os"
	"time"
)

var parentDir = "./bank/questions/"

func GenerateSimpleQuestion(dirName, quesName string) {
	dirName = parentDir + types.SimpleQues + "/q" + dirName
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		log.Fatalf("GenerateSimpleQuestion error: %v\n", err)
	}

	CreateQues(dirName, quesName)
}

func GenerateMediumQuestion(dirName, quesName string) {
	dirName = parentDir + types.MediumQues + "/q" + dirName
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		log.Fatalf("GenerateMediumQuestion error: %v\n", err)
	}

	CreateQues(dirName, quesName)
}

func GenerateHardQuestion(dirName, quesName string) {
	dirName = parentDir + types.HardQues + "/q" + dirName
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		log.Fatalf("GenerateHardQuestion error: %v\n", err)
	}

	CreateQues(dirName, quesName)
}

func CreateQues(dirName, quesName string) {
	dirName += "/" + quesName + ".go"
	file, err := os.Create(dirName)
	if err != nil {
		log.Fatalf("CreateQues error: %v\n", err)
	}
	defer file.Close()

	data := fmt.Sprintf("package main\n /**\n * 【】\n * 🪢\n * @%s\n */", time.Now().Format("2006-01-02 15:04"))
	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln(err)
	}
}

//func AddFileToGit(fileName string) {
//	cmd := exec.Command("git add", fileName)
//	if err := cmd.Run(); err != nil {
//		log.Fatalf("AddFileToGit error: %v\n", err)
//	}
//
//}
