package verbose

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func tagf(tag string, status string, color *color.Color) (int, error) {
	t := time.Now()
	return fmt.Printf(
		"%s| %02d.%02d.%d - %02d:%02d:%02d |%s| ",
		color.Sprintf(" %s%s ", tag, strings.Repeat(" ", 15-len(tag))),
		t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second(),
		color.Sprintf(" %-5s ", status),
	)
}

// Deprecated: This function has been replaced by Debug
func Print(tag string, a ...string) (n int, err error) {
	a = append(a, "\n")
	return Printf(tag, strings.Join(a, ""))
}

// Deprecated: This function has been replaced by Debugf
func Printf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Print", neutralColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}

func Debug(tag string, a ...string) (n int, err error) {
	a = append(a, "\n")
	return Debugf(tag, strings.Join(a, ""))
}
func Debugf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Debug", debugColor); err != nil {
		return 0, err
	}
	return fmt.Fprintf(os.Stdout, format, a...)
}

func Warn(tag string, a ...string) (int, error) {
	a = append(a, "\n")
	return Warnf(tag, strings.Join(a, ""))
}
func Warnf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Warn", warnColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}

func Info(tag string, a ...string) (int, error) {
	a = append(a, "\n")
	return Infof(tag, strings.Join(a, ""))
}
func Infof(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Info", infoColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}

func Success(tag string, a ...string) (n int, err error) {
	a = append(a, "\n")
	return Successf(tag, strings.Join(a, ""))
}
func Successf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Ok", successColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}

func Error(tag string, a ...string) (int, error) {
	a = append(a, "\n")
	return Errorf(tag, strings.Join(a, ""))
}
func Errorf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Error", errorColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}
