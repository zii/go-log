package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Level int

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
	file     string
	line     int
	datetime string
}

type Logger struct {
	root string // project root path
	out  io.Writer
	lv   Level
}

func New(out io.Writer, root_depth int, lv Level) *Logger {
	if out == nil {
		out = os.Stdout
	}
	if root_depth < 0 {
		root_depth = 0
	}
	_, pwd, _, _ := runtime.Caller(1)
	for i := 0; i < root_depth+1; i++ {
		pwd = filepath.Dir(pwd)
	}
	l := &Logger{
		root: pwd + "/",
		out:  out,
		lv:   lv,
	}
	return l
}

func (l *Logger) SetRoot(root_depth int) {
	_, pwd, _, _ := runtime.Caller(1)
	for i := 0; i < root_depth+1; i++ {
		pwd = filepath.Dir(pwd)
	}
	l.root = pwd + "/"
}

func (l *Logger) SetOutput(w io.Writer) {
	l.out = w
}

func (l *Logger) newHeader(skip int) *Header {
	_, file, line, _ := runtime.Caller(skip)
	if l.root != "" && strings.HasPrefix(file, l.root) {
		file = file[len(l.root):]
	}
	h := &Header{
		file:     file,
		line:     line,
		datetime: time.Now().Format(time.DateTime),
	}
	return h
}

func (l *Logger) Writeln(lv Level, tag string, v ...interface{}) {
	if l.out == nil || lv > l.lv && l.lv > LvNone {
		return
	}
	h := l.newHeader(3)
	fmt.Fprintf(l.out, "%s %s %s:%d %v\n", h.datetime, tag, h.file, h.line, fmt.Sprint(v...))
}

func (l *Logger) Writef(lv Level, tag string, format string, v ...interface{}) {
	if l.out == nil || lv > l.lv {
		return
	}
	h := l.newHeader(3)
	fmt.Fprintf(l.out, "%s %s:%d %s: %v\n", h.datetime, h.file, h.line, tag, fmt.Sprintf(format, v...))
}

func (l *Logger) Trace(v ...interface{}) {
	l.Writeln(LvTrace, "TRC", v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Writef(LvTrace, "TRC", format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.Writeln(LvError, "ERR", v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Writef(LvError, "ERR", format, v...)
}

var Default = New(os.Stderr, 0, LvInfo)

func Error(v ...interface{}) {
	Default.Writeln(LvError, "ERR", v...)
}
