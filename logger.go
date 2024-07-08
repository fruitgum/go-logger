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
	logOut      = os.Stdout
	logPrint    = log.New(logOut, "", log.Ldate|log.Ltime)
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
		System("Redirect to %v", filepath.Join(dir, logfile))
		logOut = writeTo
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

	writeTo, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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

func writeLog(logColor, prefixStr, message string) {
	formatedPrefixStr := "[" + prefixStr + "]"
	var prefix string
	switch logColor {
	case "green":
		prefix = color.GreenString(formatedPrefixStr)
	case "yellow":
		prefix = color.YellowString(formatedPrefixStr)
	case "red":
		prefix = color.RedString(formatedPrefixStr)
	case "blue":
		prefix = color.BlueString(formatedPrefixStr)
	case "magenta":
		prefix = color.MagentaString(formatedPrefixStr)
	case "cyan":
		prefix = color.CyanString(formatedPrefixStr)
	case "white":
		prefix = color.WhiteString(formatedPrefixStr)
	}
	if logOut == os.Stdout {
		logPrint.Println(fmt.Sprintf("%s %s", prefix, message))
	} else {
		logPrint.Println(fmt.Sprintf("%s %s", formatedPrefixStr, message))
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
		//logPrint.Println(fmt.Sprintf("%s %s", color.GreenString("[OK]"), message))
		writeLog("green", "OK", message)
	}
}

// System Printing message with format 01/01/1970 00:00:00 [SYS] string
//
// Suppressing if log level set to none or fatal
func System(format string, v ...interface{}) {
	if LogLevelInt >= 0 {
		message := fmt.Sprintf(format, v...)
		//logPrint.Println(fmt.Sprintf("%s %s", color.WhiteString("[SYS]"), message))
		writeLog("white", "SYS", message)
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
		//logPrint.Println(fmt.Sprintf("%s %s", color.RedString("[FATAL]"), message))
		writeLog("red", "FATAL", message)
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
		//logPrint.Println(fmt.Sprintf("%s %s", color.RedString("[ERROR]"), message))
		writeLog("red", "ERROR", message)
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
		//logPrint.Println(fmt.Sprintf("%s %s", color.YellowString("[WARN]"), message))
		writeLog("yellow", "WARN", message)
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
		//logPrint.Println(fmt.Sprintf("%s %s", color.CyanString("[INFO]"), message))
		writeLog("cyan", "INFO", message)
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
		//logPrint.Println(fmt.Sprintf("%s %s", color.MagentaString("[DEBUG]"), message))
		writeLog("magenta", "DEBUG", message)
	}
}
