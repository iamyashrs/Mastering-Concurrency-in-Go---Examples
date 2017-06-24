package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/alecthomas/log4go"
)

type LogItem struct {
	Message string
}

var Logs []LogItem

func saveLogs() {
	logFile := log4go.NewFileLogWriter("stack.log", false)
	logFile.SetFormat("%d %t - %M (%S)")
	logFile.SetRotate(false)
	logFile.SetRotateSize(0)
	logFile.SetRotateDaily(true)

	logStack := make(log4go.Logger)
	logStack.AddFilter("file", log4go.DEBUG, logFile)
	for i := range Logs {
		fmt.Println(Logs[i].Message)
		logStack.Info(Logs[i].Message)
	}
}

func goDetails(done chan bool) {
	i := 0
	for {
		var message string
		stackBuf := make([]byte, 1024)
		stack := runtime.Stack(stackBuf, false)
		stack++

		_, callerFile, callerLine, ok := runtime.Caller(0)
		message = "Goroutinge from " + string(callerLine) + "" + string(callerFile) + " stack:" + string(stackBuf)

		openGoroutine := runtime.NumGoroutine()

		if ok == true {
			message = message + callerFile
		}

		message = message + strconv.FormatInt(int64(openGoroutine), 10) + " goroutine active"

		li := LogItem{Message: message}

		Logs = append(Logs, li)

		if i == 20 {
			done <- true
			break
		}
		i++
	}
}

func main() {
	done := make(chan bool)

	go goDetails(done)
	for i := 0; i < 10; i++ {
		go goDetails(done)
	}

	for {
		select {
		case d := <-done:
			if d == true {
				saveLogs()
				os.Exit(1)
			}
		}
	}
}
