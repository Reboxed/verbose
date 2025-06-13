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
	return Printf(tag, "%s\n", strings.Join(a, ""))
}

// Deprecated: This function has been replaced by Debugf
func Printf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Print", NeutralColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}

func Debug(tag string, a ...string) (n int, err error) {
	return Debugf(tag, "%s\n", strings.Join(a, ""))
}
func Debugf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Debug", DebugColor); err != nil {
		return 0, err
	}
	return fmt.Fprintf(os.Stdout, format, a...)
}

func Custom(tag string, typeStr string, color *color.Color, a ...string) (n int, err error) {
	return Customf(tag, typeStr, color, "%s\n", strings.Join(a, ""))
}
func Customf(tag string, typeStr string, color *color.Color, format string, a ...interface{}) (n int, err error) {
	if _, err := tagf(tag, typeStr, color); err != nil {
		return 0, err
	}
	return fmt.Fprintf(os.Stdout, format, a...)
}

func Warn(tag string, a ...string) (int, error) {
	return Warnf(tag, "%s\n", strings.Join(a, ""))
}
func Warnf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Warn", WarnColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}

func Info(tag string, a ...string) (int, error) {
	return Infof(tag, "%s\n", strings.Join(a, ""))
}
func Infof(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Info", InfoColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}

func Success(tag string, a ...string) (n int, err error) {
	return Successf(tag, "%s\n", strings.Join(a, ""))
}
func Successf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Ok", SuccessColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}

func Error(tag string, a ...string) (int, error) {
	return Errorf(tag, "%s\n", strings.Join(a, ""))
}
func Errorf(tag string, format string, a ...interface{}) (int, error) {
	if _, err := tagf(tag, "Error", ErrorColor); err != nil {
		return 0, err
	}

	return fmt.Fprintf(os.Stdout, format, a...)
}
