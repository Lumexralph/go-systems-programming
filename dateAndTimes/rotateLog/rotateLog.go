package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	TOTALWRITES       = 0
	ENTRIESPERLOGFILE = 100
	WHENTOSTOP        = 230
	openLogFile       os.File
)

func rotateLogFile(filename string) error {
	openLogFile.Close()
	err := os.Rename(filename, filename+"."+strconv.Itoa(TOTALWRITES))
	if err != nil {
		return err
	}
	err = setUpLogFile(filename)
	return err
}

func setUpLogFile(filename string) error {
	openLogFile, err := os.OpenFile(filename,
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(openLogFile)
	return nil
}

func main() {
	numberOfLogEntries := 0
	ipAddress := ""
	filename := "/tmp/myLog.log"
	err := setUpLogFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for {
		if numberOfLogEntries%10 != 0 {
			ipAddress = "127.0.0.1"
		} else {
			ipAddress = ""
		}
		log.Printf("%d: %s This is a test log entry  \n",
			numberOfLogEntries, ipAddress)
		numberOfLogEntries++
		TOTALWRITES++

		if numberOfLogEntries > ENTRIESPERLOGFILE {
			rotateLogFile(filename)
			numberOfLogEntries = 0
		}

		if TOTALWRITES > WHENTOSTOP {
			rotateLogFile(filename)
			break
		}
		time.Sleep(time.Second)
	}

	fmt.Printf("Wrote %d log entries! %s\n", TOTALWRITES, ipAddress)
}
