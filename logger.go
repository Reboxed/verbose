package verbose

import (
    "fmt"
    "os"
    "strings"
    "time"

    "github.com/fatih/color"
)

func Debugf(tag string, format string, status string, color *color.Color, a []string, b ...interface{}) (n int, err error) {
    t := time.Now()
    _, tagerr := fmt.Printf("%s| %02d.%02d.%d - %02d:%02d:%02d |%s| ",
        color.Sprintf(" %s%s ", tag, strings.Repeat(" ", 15-len(tag))),
        t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second(),
        color.Sprintf(" %-5s ", status))

    if tagerr != nil {
        return 0, tagerr
    }

    return fmt.Fprintf(os.Stdout, strings.Join(append(a, "\n"), ""), b...)
}