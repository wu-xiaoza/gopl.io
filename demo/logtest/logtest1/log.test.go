package main

import (
	"fmt"
	"log"
)

func main() {
	log.Print("qwe\n")
	log.Printf("qwe\n")
	log.Println("qwe")
}

func test() {
	fmt.Println("first fatal defer")
	fmt.Println("sceond fatal defer")
}

func test1() {
	fmt.Println("first panic defer")
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	defer test1()
	//Panic is equivalent to Print() followed by a call to panic().
	log.Panic("this is log panic test\n")
	defer func() {
		fmt.Println("sceond panic test")
	}()
}

// func init() {
// 	defer test()
// 	//Fatal is equivalent to Print() followed by a call to os.Exit(1).
// 	//os.Exit(1)执行后defer没有作用
// 	log.Fatal("this is first log fatal test\n")
// 	//Output writes the output for a logging event. The string s contains the text to print after the prefix specified by the flags of the Logger. A newline is appended if the last character of s is not already a newline. Calldepth is used to recover the PC and is provided for generality, although at the moment on all pre-defined paths it will be 2.
// 	/* func Fatal(v ...interface{}) {
// 		std.Output(2, fmt.Sprint(v...))
// 		os.Exit(1)
// 	} */
// 	//os.Exit(1)后面不会执行
// 	log.Fatalf("this is sceond log fatal test\n")
// 	log.Fatalln("this is third log fatal test")
// }


//Caller reports file and line number information about function invocations on the calling goroutine's stack. The argument skip is the number of stack frames to ascend, with 0 identifying the caller of Caller. (For historical reasons the meaning of skip differs between Caller and Callers.) The return values report the program counter, file name, and line number within the file of the corresponding call. The boolean ok is false if it was not possible to recover the information.

//for accumulating text to write