package log

import (
	"fmt"
	"github.com/zii/go-log/color"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	LvNone = iota
	LvFatal
	LvError
	LvWarn
	LvInfo
	LvDebug
	LvTrace

	LvMax
)

type Header struct {
	file string
	line int
	time string
}

type Logger struct {
	root        string // project base dir
	out         io.Writer
	lv          int
	time_format string
}

func New(out io.Writer, up_levels int, lv int) *Logger {
	if out == nil {
		out = os.Stderr
	}
	l := &Logger{
		out:         out,
		lv:          lv,
		time_format: time.DateTime,
	}
	if up_levels >= 0 {
		_, pwd, _, _ := runtime.Caller(1)
		for i := 0; i < up_levels+1; i++ {
			pwd = filepath.Dir(pwd)
		}
		l.root = pwd + string(filepath.Separator)
	}
	return l
}

// SetRoot Used to reduce file path length in log printing
// up_levels: how many up levels to project base dir
func (l *Logger) SetRoot(up_levels int) {
	if up_levels < 0 {
		l.root = ""
		return
	}
	_, pwd, _, _ := runtime.Caller(1)
	for i := 0; i < up_levels+1; i++ {
		pwd = filepath.Dir(pwd)
	}
	l.root = pwd + string(filepath.Separator)
}

func (l *Logger) Root() string {
	return l.root
}

func (l *Logger) SetLevel(lv int) {
	l.lv = lv
}

func (l *Logger) Level() int {
	return l.lv
}

func (l *Logger) SetOutput(w io.Writer) {
	l.out = w
}

func (l *Logger) SetTimeFormat(f string) {
	l.time_format = f
}

func (l *Logger) newHeader(skip int) *Header {
	_, file, line, _ := runtime.Caller(skip)
	if l.root != "" && strings.HasPrefix(file, l.root) {
		file = file[len(l.root):]
	}
	h := &Header{
		file: file,
		line: line,
	}
	if l.time_format != "" {
		h.time = time.Now().Format(l.time_format)
	}
	return h
}

func ColorString(lv int, s string) string {
	var c string
	switch lv {
	case LvTrace:
		c = color.FgBlue
	case LvDebug:
		c = color.FgCyan
	case LvInfo:
		c = color.FgGreen
	case LvWarn:
		c = color.FgYellow
	case LvError:
		c = color.FgRed
	case LvFatal:
		c = color.FgMagenta
	}
	return color.String(s, c)
}

func (l *Logger) Writeln(lv int, tag string, v ...interface{}) {
	if l.out == nil || lv > l.lv && l.lv > LvNone {
		return
	}
	h := l.newHeader(3)
	tag = ColorString(lv, tag)
	fmt.Fprintf(l.out, "%s %s %s:%d %s", h.time, tag, h.file, h.line, fmt.Sprintln(v...))
}

func (l *Logger) Writef(lv int, tag string, format string, v ...interface{}) {
	if l.out == nil || lv > l.lv && l.lv > LvNone {
		return
	}
	h := l.newHeader(3)
	tag = ColorString(lv, tag)
	fmt.Fprintf(l.out, "%s %s %s:%d %s\n", h.time, tag, h.file, h.line, fmt.Sprintf(format, v...))
}

func (l *Logger) Trace(v ...interface{}) {
	l.Writeln(LvTrace, "T", v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Writef(LvTrace, "T", format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Writeln(LvDebug, "D", v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Writef(LvDebug, "D", format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.Writeln(LvInfo, "I", v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Writef(LvInfo, "I", format, v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.Writeln(LvInfo, "I", v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.Writef(LvInfo, "I", format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.Writeln(LvWarn, "W", v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Writef(LvWarn, "W", format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.Writeln(LvError, "E", v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Writef(LvError, "E", format, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Writeln(LvFatal, "F", v...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Writef(LvFatal, "F", format, v...)
	os.Exit(1)
}

var Default = New(os.Stderr, -1, LvInfo)

var SetRoot = Default.SetRoot
var Root = Default.Root
var SetLevel = Default.SetLevel
var Level = Default.Level
var SetTimeFormat = Default.SetTimeFormat

var Trace = Default.Trace
var Tracef = Default.Tracef
var Debug = Default.Debug
var Debugf = Default.Debugf
var Info = Default.Info
var Infof = Default.Infof
var Println = Default.Println
var Printf = Default.Printf
var Warn = Default.Warn
var Warnf = Default.Warnf
var Error = Default.Error
var Errorf = Default.Errorf
var Fatal = Default.Fatal
var Fatalf = Default.Fatalf
