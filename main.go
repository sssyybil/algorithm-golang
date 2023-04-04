package main

import (
	"algorithm-golang/utils"
	"flag"
	"log"
)

func main() {
	dirName := flag.String("dName", "000", "Please type question number.")
	qName := flag.String("qName", "bar", "Please type question name.")
	qStatus := flag.String("status", "000", "s=simple; m=medium; h=hard")
	flag.Parse()

	switch *qStatus {
	case "s":
		utils.GenerateSimpleQuestion(*dirName, *qName)
	case "m":
		utils.GenerateMediumQuestion(*dirName, *qName)
	case "h":
		utils.GenerateHardQuestion(*dirName, *qName)
	default:
		log.Fatalln("Status is not exists.")
	}
}
