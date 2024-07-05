package logger

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

var logLevelMap = map[string]int{
	"debug":   4,
	"info":    3,
	"warn":    2,
	"fatal":   1,
	"error":   1,
	"minimal": 0,
	"none":    -1,
}

var (
	LogLevelInt = 3
	logPrint    = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

func SetLogLevel(logLevel string) string {
	if level, exists := logLevelMap[logLevel]; exists {
		LogLevelInt = level
		return logLevel
	} else {
		System("Invalid log level %v", logLevel)
		LogLevelInt = logLevelMap["info"]
		return "info"
	}

}

func HelpUsage() string {
	loggerUsage :=
		"Available levels:\n" +
			"debug\n" +
			"info\n" +
			"warn\n" +
			"error\n" +
			"minimal\n" +
			"none"
	return loggerUsage
}

func Success(format string, v ...interface{}) {
	if LogLevelInt >= 0 {
		message := fmt.Sprintf(format, v...)
		logPrint.Println(fmt.Sprintf("%s %s", color.GreenString("[OK]"), message))
	}
}

func System(format string, v ...interface{}) {
	if LogLevelInt >= 0 {
		message := fmt.Sprintf(format, v...)
		logPrint.Println(fmt.Sprintf("%s %s", color.WhiteString("[SYS]"), message))
	}
}

func Fatal(format string, v ...interface{}) {
	if LogLevelInt > -1 {
		message := fmt.Sprintf(format, v...)
		logPrint.Println(fmt.Sprintf("%s %s", color.RedString("[ERROR]"), message))
	}
	os.Exit(1)
}

func Error(format string, v ...interface{}) {
	if LogLevelInt > 0 {
		message := fmt.Sprintf(format, v...)
		logPrint.Println(fmt.Sprintf("%s %s", color.RedString("[ERROR]"), message))
	}
}

func Warn(format string, v ...interface{}) {
	if LogLevelInt > 1 {
		message := fmt.Sprintf(format, v...)
		logPrint.Println(fmt.Sprintf("%s %s", color.YellowString("[WARN]"), message))
	}
}

func Info(format string, v ...interface{}) {
	if LogLevelInt > 2 {
		message := fmt.Sprintf(format, v...)
		logPrint.Println(fmt.Sprintf("%s %s", color.CyanString("[INFO]"), message))
	}
}

func Debug(format string, v ...interface{}) {
	if LogLevelInt > 3 {
		message := fmt.Sprintf(format, v...)
		logPrint.Println(fmt.Sprintf("%s %s", color.MagentaString("[DEBUG]"), message))
	}
}
