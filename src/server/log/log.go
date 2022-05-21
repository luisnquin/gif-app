package log

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

const (
	FATAL string = "FATAL"
	ERROR string = "ERROR"
	WARN  string = "WARN"
	INFO  string = "INFO"
)

func Error(err error) {
	color.New(color.FgHiRed).Println(time.Now().Format("2006-01-02 15:04:05") + " " + ERROR + " " + err.Error())
}

func Errorf(format string, args ...any) {
	capsule := fmt.Errorf(format, args...)
	color.New(color.FgHiRed).Println(time.Now().Format("2006-01-02 15:04:05") + " " + ERROR + " " + capsule.Error())
}

func Info(message ...any) {
	capsule := fmt.Sprint(message...)
	color.New(color.FgHiBlue).Println(time.Now().Format("2006-01-02 15:04:05") + " " + INFO + " " + capsule)
}

func Infof(format string, args ...any) {
	capsule := fmt.Sprintf(format, args...)
	color.New(color.FgHiBlue).Println(time.Now().Format("2006-01-02 15:04:05") + " " + INFO + " " + capsule)
}

func Fatal(message ...any) {
	capsule := fmt.Sprint(message...)
	color.New(color.FgMagenta).Println(time.Now().Format("2006-01-02 15:04:05") + " " + FATAL + " " + capsule)
	os.Exit(1)
}

func Fatalf(format string, args ...any) {
	capsule := fmt.Sprintf(format, args...)
	color.New(color.FgMagenta).Println(time.Now().Format("2006-01-02 15:04:05") + " " + FATAL + " " + capsule)
	os.Exit(1)
}

func FatalWithCheck(err error) {
	if err != nil {
		Fatal(err)
	}
}

func Warn(message ...any) {
	capsule := fmt.Sprint(message...)
	color.New(color.FgHiYellow).Println(time.Now().Format("2006-01-02 15:04:05") + " " + WARN + " " + capsule)
}

func Warnf(format string, args ...any) {
	capsule := fmt.Sprintf(format, args...)
	color.New(color.FgHiYellow).Println(time.Now().Format("2006-01-02 15:04:05") + " " + WARN + " " + capsule)
}
