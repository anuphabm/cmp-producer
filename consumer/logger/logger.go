package logger

import (
	"consumer/config"
	"fmt"
	"os"
	"strings"

	logger "github.com/sirupsen/logrus"
)

func setLogLevel(logLevel string) {
	switch strings.ToLower(logLevel) {
	case "debug":
		logger.SetLevel(logger.DebugLevel)
	case "info":
		logger.SetLevel(logger.InfoLevel)
	case "warn":
		logger.SetLevel(logger.WarnLevel)
	case "error":
		logger.SetLevel(logger.ErrorLevel)
	default:
		logger.SetLevel(logger.DebugLevel)
	}
}

func Init() {
	customFormatter := new(logger.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logger.SetFormatter(customFormatter)
	logger.SetReportCaller(true)
	envMode := os.Getenv("RUN_MODE")

	logLevel := config.Appconfig.GetString(fmt.Sprintf("%s.Logging.level", envMode))
	setLogLevel(logLevel)
	if config.Appconfig.GetBool(fmt.Sprintf("%s.Logging.stdout", envMode)) {
		logger.New().Out = os.Stdout
	} else {
		file, err := os.OpenFile(config.Appconfig.GetString("Logging.path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		} else {
			fmt.Println("Failed to log to file ", err.Error())
		}
	}

}

// Panic will exit with status code 2
func PanicLn(message string) {
	logger.Panicln(message)
}

// Fatal will exit with status code 1
func FatalLn(message string) {
	logger.Fatalln(message)
}

// Just log the message as info
func InfoLn(message string) {
	logger.Infoln(message)
}

// Just log the message as warn
func WarnLn(message string) {
	logger.Warnln(message)
}

// Just log the message as debug
func DebugLn(message string) {
	logger.Debugln(message)
}

// Just log the message as debug
func PrintLn(message string) {
	logger.Print(message)
}
