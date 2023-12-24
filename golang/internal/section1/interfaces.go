package section1

import (
	"fmt"
	"os"
	"time"
)

// Logger is the interface that defines the Log method.
type Logger interface {
	Log(message string)
}

// FileLogger is a struct that implements the Logger interface to log messages to a file.
type FileLogger struct {
	FileName string
}

// ConsoleLogger is a struct that implements the Logger interface to log messages to the console.
type ConsoleLogger struct{}

// Log writes the message to the console.
func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

// Log writes the message to a file.
func (f FileLogger) Log(message string) {
	file, err := os.OpenFile(f.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	logMessage := fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), message)

	if _, err := file.WriteString(logMessage); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
