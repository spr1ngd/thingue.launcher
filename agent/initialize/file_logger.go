package initialize

import (
	"github.com/wailsapp/wails/v2/pkg/logger"
	"log"
	"os"
)

// FileLogger is a utility to log messages to a number of destinations
type FileLogger struct {
	filename string
}

// NewFileLogger creates a new Logger.
func NewFileLogger(filename string) logger.Logger {
	return &FileLogger{
		filename: filename,
	}
}

// Print works like Sprintf.
func (l *FileLogger) Print(message string) {
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	myLog := log.New(f, "wails_", log.Ldate|log.Ltime|log.Lshortfile)
	//设置输出前缀
	myLog.SetPrefix("wails_")
	//输出一条日志
	myLog.Print(message)
	//if _, err := f.WriteString(message); err != nil {
	//	f.Disconnect()
	//	log.Fatal(err)
	//}
	//f.Disconnect()
}

func (l *FileLogger) Println(message string) {
	l.Print(message + "\n")
}

// Trace level logging. Works like Sprintf.
func (l *FileLogger) Trace(message string) {
	l.Println("TRACE | " + message)
}

// Debug level logging. Works like Sprintf.
func (l *FileLogger) Debug(message string) {
	l.Println("DEBUG | " + message)
}

// Info level logging. Works like Sprintf.
func (l *FileLogger) Info(message string) {
	l.Println("INFO  | " + message)
}

// Warning level logging. Works like Sprintf.
func (l *FileLogger) Warning(message string) {
	l.Println("WARN  | " + message)
}

// Error level logging. Works like Sprintf.
func (l *FileLogger) Error(message string) {
	l.Println("ERROR | " + message)
}

// Fatal level logging. Works like Sprintf.
func (l *FileLogger) Fatal(message string) {
	l.Println("FATAL | " + message)
	os.Exit(1)
}
