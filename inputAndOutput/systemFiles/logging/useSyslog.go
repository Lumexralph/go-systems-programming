package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])

	// on MacOs the message of the syslog is sent to the system.log file
	// in the path /var/log/system.log
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Crit("Crit: Logging in Go!")

	sysLog, err = syslog.New(syslog.LOG_ALERT|syslog.LOG_LOCAL7, "Some Program")
	if err != nil {
		log.Fatal(sysLog, err)
	}
	sysLog.Emerg("Emergency priority: Logging in Go!")

	fmt.Fprint(sysLog, "Log.Print: Logging in Go!")
}
