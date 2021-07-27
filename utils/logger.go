package utils

import (
	"fmt"
	"time"
)

const(
	LogInfo = "INFO"
	LogWarning = "WARNING"
	LogError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func Log(message string, severity string) {
	logCh <- logEntry{ time.Now(), severity, message}
}

func EndLog(){
	doneCh <- struct{}{}
}

func InitLogger(){
	// for entry := range logCh {
	// 	fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	// }

	for {
		select {
		case entry := <- logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <- doneCh:
			break
		}
	}
}
