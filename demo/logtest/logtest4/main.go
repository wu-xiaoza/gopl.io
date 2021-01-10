package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//OpenFile is the generalized open call; most users will use Open or Create instead. It opens the named file with specified flag (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag is passed, it is created with mode perm (before umask). If successful, methods on the returned File can be used for I/O. If there is an error, it will be of type *PathError.
	//openFileNolog is the Windows implementation of OpenFile.
	//whether file is opened for appending
	logFile, err := os.OpenFile("./demo/logtest/logtest4/xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")

	var logger = log.New(os.Stdout, "", log.LstdFlags)
	logger.SetFlags(log.Lshortfile)
	logger.SetPrefix("aaa")
	logger.Println("hello, world")
}
