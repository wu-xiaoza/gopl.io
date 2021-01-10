package main

import (
	"fmt"
	"log"
	"os"
)

//Debug .
func Debug(logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create ./test.log err : %v\n", err)
	}
	if logFile != nil {
		defer func(file *os.File) { file.Close() }(logFile)
	}

	debugLog := log.New(logFile, "[Debug]", log.Ldate)

	debugLog.SetPrefix("[Debug]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println("this is Debug log")
}

//Waring .
func Waring(logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create ./test.log err : %v\n", err)
	}
	if logFile != nil {
		defer func(file *os.File) { file.Close() }(logFile)
	}

	debugLog := log.New(logFile, "[Waring]", log.Ldate)

	debugLog.SetPrefix("[Waring]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println("this is Waring log")
}

func main() {
	logName := "./demo/logtest/logtest2/test.log"
	Debug(logName)
	Waring(logName)
}
