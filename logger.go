package logger

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"path/filepath"
	"strings"
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

var (
	LogLevelInt = 3
	logToFile   = false
	logPrint    = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

// SetLogLevel accepts one of the following strings:
//
//   - debug
//   - info
//   - warn
//   - error
//   - fatal
//   - none
//
// and sets log level according to given parameter
//
// # Returns given parameter sting
//
// logLevel default value is info
func SetLogLevel(logLevel string) string {
	logLevel = strings.ToLower(logLevel)
	if level, exists := logLevelMap[logLevel]; exists {
		LogLevelInt = level
		return logLevel
	} else {
		System("Invalid log level %v", logLevel)
		LogLevelInt = logLevelMap["info"]
		return "info"
	}

}

// ToFile redirects output to defined file
//
//	logger.ToFile("dir", "file")
//
// If neither dir nor file is set, the logger will create a default file ./logs/*current_date*.log
func ToFile(dir, logfile string) {

	writeTo, err := checkLogFile(dir, logfile)
	if err != nil {
		Error("Can not open log file %v: %v", filepath.Join(dir, logfile), err)
	} else {
		System("Redirect to %v", writeTo.Name())
		logToFile = true
		logPrint = log.New(writeTo, "", log.Ldate|log.Ltime)
	}
}

func checkLogFile(dir, logfile string) (*os.File, error) {
	if dir == "" {
		dir = "logs"
	}

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	if logfile == "" {
		logfile = time.Now().Format("2006-01-02") + ".log"
	}

	file := filepath.Join(dir, logfile)

	writeTo, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return writeTo, nil

}

func HelpUsage() string {
	loggerUsage := "Case-insensitive\n" +
		"Available levels:\n" +
		"debug\n" +
		"info\n" +
		"warn\n" +
		"error\n" +
		"fatal\n" +
		"none"
	return loggerUsage
}

func writeLog(logColor, prefixStr, message string) string {
	formatedPrefixStr := "[" + prefixStr + "]"
	if !logToFile {
		var prefix string
		switch logColor {
		case "green":
			prefix = fmt.Sprintf(color.GreenString(formatedPrefixStr))
		case "yellow":
			prefix = fmt.Sprintf(color.YellowString(formatedPrefixStr))
		case "red":
			prefix = fmt.Sprintf(color.RedString(formatedPrefixStr))
		case "blue":
			prefix = fmt.Sprintf(color.BlueString(formatedPrefixStr))
		case "magenta":
			prefix = fmt.Sprintf(color.MagentaString(formatedPrefixStr))
		case "cyan":
			prefix = fmt.Sprintf(color.CyanString(formatedPrefixStr))
		case "white":
			prefix = fmt.Sprintf(color.WhiteString(formatedPrefixStr))
		}
		return prefix
	} else {
		return formatedPrefixStr
	}

}

// Success Printing message with format
//
//	01/01/1970 00:00:00 [OK] string
//
// Suppressing if log level set to none or fatal
func Success(format string, v ...interface{}) {
	if LogLevelInt >= 0 {
		message := fmt.Sprintf(format, v...)
		formatted := writeLog("green", "OK", message)
		logPrint.Println(fmt.Sprintf("%s %s", formatted, message))
	}
}

// System Printing message with format 01/01/1970 00:00:00 [SYS] string
//
// Suppressing if log level set to none or fatal
func System(format string, v ...interface{}) {
	if LogLevelInt >= 0 {
		message := fmt.Sprintf(format, v...)
		formatted := writeLog("white", "SYS", message)
		logPrint.Println(fmt.Sprintf("%s %s", formatted, message))
	}
}

// Fatal Printing message with format
//
//	01/01/1970 00:00:00 [FATAL] string
//
// # Suppressing if log level set to none
//
// Will terminate process even if its output is suppressed
func Fatal(format string, v ...interface{}) {
	if LogLevelInt > -1 {
		message := fmt.Sprintf(format, v...)
		formatted := writeLog("red", "FATAL", message)
		logPrint.Println(fmt.Sprintf("%s %s", formatted, message))
	}
	os.Exit(1)
}

// Error Printing message with format
//
//	01/01/1970 00:00:00 [ERROR] string
//
// # Suppressing if log level set to fatal or none
//
// Won't terminate process
func Error(format string, v ...interface{}) {
	if LogLevelInt > 0 {
		message := fmt.Sprintf(format, v...)
		formatted := writeLog("red", "ERROR", message)
		logPrint.Println(fmt.Sprintf("%s %s", formatted, message))
	}
}

// Warn Printing message with format
//
//	#### 01/01/1970 00:00:00 [WARN] string
//
// Suppressing if log level set to error, fatal or none
func Warn(format string, v ...interface{}) {
	if LogLevelInt > 1 {
		message := fmt.Sprintf(format, v...)
		formatted := writeLog("yellow", "WARN", message)
		logPrint.Println(fmt.Sprintf("%s %s", formatted, message))
	}
}

// Info Printing message with format
//
//	01/01/1970 00:00:00 [INFO] string
//
// Suppressing if log level set to warn, error, fatal or none
func Info(format string, v ...interface{}) {
	if LogLevelInt > 2 {
		message := fmt.Sprintf(format, v...)
		formatted := writeLog("cyan", "INFO", message)
		logPrint.Println(fmt.Sprintf("%s %s", formatted, message))
	}
}

// Debug Printing message with format
//
//	01/01/1970 00:00:00 [DEBUG] string
//
// Suppressing if log level is not set to debug
func Debug(format string, v ...interface{}) {
	if LogLevelInt > 3 {
		message := fmt.Sprintf(format, v...)
		formatted := writeLog("magenta", "DEBUG", message)
		logPrint.Println(fmt.Sprintf("%s %s", formatted, message))
	}
}
