package goLogger

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var logLevelMap = map[string]int{
	"debug": 4,
	"info":  3,
	"warn":  2,
	"error": 1,
	"fatal": 0,
	"none":  -1,
}

type Logger struct {
	logLevel  int
	LogOutput io.Writer
	LogToFile bool
}

func New() *Logger {
	var logger Logger
	logger.logLevel = 3
	logger.LogToFile = false
	logger.LogOutput = os.Stdout
	return &logger
}

// SetLogLevel
//
// Available levels
//
//	debug - will show everything
//	info  - drops debug messages
//	warn  - drops info and debug messages
//	error - drops warn, info and debug messages
//	fatal - drops error, warn, info and debug messages.
//	none  - will show nothing
func (s *Logger) SetLogLevel(level string) {
	s.logLevel = logLevelMap[level]
}

// ToFile writes logs into file
//
// If dir is not set ("") - it will create "logs" folder next to an app
func (s *Logger) ToFile(dir string) {
	subFolderMonth := strconv.Itoa(int(time.Now().Month()))
	subFolderYear := strconv.Itoa(time.Now().Year())

	if len(subFolderMonth) == 1 {
		subFolderMonth = "0" + subFolderMonth
	}

	if dir == "" {
		dir = "logs"
	}

	dir = filepath.Join(dir, subFolderYear, subFolderMonth)
	_, err := os.ReadDir(dir)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("dir:", err)
		}
	}

	logFileName := time.Now().Format("2006-01-02") + ".log"
	logFullPath := filepath.Join(dir, logFileName)
	file, err := os.OpenFile(logFullPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("file:", err)
	}
	s.LogToFile = true
	s.LogOutput = file

}

func LogMessage(logPrefix, logMessage string, s *Logger, messageLogLevel int) {

	if messageLogLevel > s.logLevel {
		return
	}

	logToFile := s.LogToFile
	logOutput := s.LogOutput

	logMessage = fmt.Sprintf("%s   %s", logPrefix, logMessage)
	logTimeStamp := time.Now().Format("2006-01-02 15:04:05")

	logRow := logTimeStamp + " " + logMessage
	if logToFile {
		_, err := logOutput.Write([]byte(logRow + "\n"))
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	} else {
		fmt.Println(logRow)
	}

}

func (s *Logger) Debug(format string, v ...interface{}) {
	minLogLevel := 4
	color.NoColor = s.LogToFile
	logMessage := fmt.Sprintf(format, v...)
	logPrefix := fmt.Sprintf(color.MagentaString("[DBG]"))
	LogMessage(logPrefix, logMessage, s, minLogLevel)
}

func (s *Logger) Info(format string, v ...interface{}) {
	minLogLevel := 3
	color.NoColor = s.LogToFile
	logMessage := fmt.Sprintf(format, v...)
	logPrefix := fmt.Sprintf(color.CyanString("[INF]"))
	LogMessage(logPrefix, logMessage, s, minLogLevel)
}

func (s *Logger) Warn(format string, v ...interface{}) {
	minLogLevel := 2
	color.NoColor = s.LogToFile
	logMessage := fmt.Sprintf(format, v...)
	logPrefix := fmt.Sprintf(color.YellowString("[WRN]"))
	LogMessage(logPrefix, logMessage, s, minLogLevel)
}

func (s *Logger) Error(format string, v ...interface{}) {
	minLogLevel := 1
	color.NoColor = s.LogToFile
	logMessage := fmt.Sprintf(format, v...)
	logPrefix := fmt.Sprintf(color.RedString("[ERR]"))
	LogMessage(logPrefix, logMessage, s, minLogLevel)
}

// Fatal will terminate process. Even if its output suppressed
func (s *Logger) Fatal(format string, v ...interface{}) {
	minLogLevel := 0
	color.NoColor = s.LogToFile
	logMessage := fmt.Sprintf(format, v...)
	logPrefix := fmt.Sprintf(color.RedString("[FATAL]"))
	LogMessage(logPrefix, logMessage, s, minLogLevel)
	os.Exit(1)
}

func (s *Logger) Success(format string, v ...interface{}) {
	minLogLevel := 0
	color.NoColor = s.LogToFile
	logMessage := fmt.Sprintf(format, v...)
	logPrefix := fmt.Sprintf(color.GreenString("[OK ]"))
	LogMessage(logPrefix, logMessage, s, minLogLevel)
}

func (s *Logger) System(format string, v ...interface{}) {
	minLogLevel := 0
	color.NoColor = s.LogToFile
	logMessage := fmt.Sprintf(format, v...)
	logPrefix := fmt.Sprintf(color.WhiteString("[SYS]"))
	LogMessage(logPrefix, logMessage, s, minLogLevel)
}
