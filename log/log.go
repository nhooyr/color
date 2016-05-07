/*
Package log implements a simple logging package with support for highlight verbs.
It defines a Logger type with methods for formatting and printing output.

It also defines a global standard Logger that writes to standard error. Color output
will only be enabled if standard error is a terminal.
Use the helper functions Printf[p], Fatalf[p], Panicf[p], and SetOutput to access it.
*/
package log

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/nhooyr/color"
)

// Logger is a very simple logger similar to log.Logger but it supports the highlight verbs.
type Logger struct {
	mu    sync.Mutex // ensures atomic writes
	out   io.Writer  // destination for output
	color bool       // enable color output
}

// New creates a new Logger. The out argument sets the
// destination to which log data will be written.
// The color argument dictates whether color output is enabled.
func New(out io.Writer, color bool) *Logger {
	return &Logger{out: out, color: color}
}

// Printf processes the highlight verbs in format and then calls
// fmt.Fprintf to print to the underlying writer.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Fprintf(l.out, color.Run(format, l.color), v...)
}

// Printfp is the same as l.Printf but takes a prepared format struct.
func (l *Logger) Printfp(f *color.Format, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Fprintf(l.out, f.Get(l.color), v...)
}

// Print calls fmt.Fprint to print to the underlying writer.
func (l *Logger) Print(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Fprint(l.out, v...)
}

// Println calls fmt.Fprintln to print to the underlying writer.
func (l *Logger) Println(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Fprintln(l.out, v...)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.mu.Lock()
	fmt.Fprintf(l.out, color.Run(format, l.color), v...)
	os.Exit(1)
}

// Fatalfp is the same as l.Fatalf but takes a prepared format struct.
func (l *Logger) Fatalfp(f *color.Format, v ...interface{}) {
	l.mu.Lock()
	fmt.Fprintf(l.out, f.Get(l.color), v...)
	os.Exit(1)
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func (l *Logger) Fatal(v ...interface{}) {
	l.mu.Lock()
	fmt.Fprint(l.out, v...)
	os.Exit(1)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (l *Logger) Fatalln(v ...interface{}) {
	l.mu.Lock()
	fmt.Fprintln(l.out, v...)
	os.Exit(1)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	s := fmt.Sprintf(format, v...)
	io.WriteString(l.out, s)
	panic(s)
}

// Panicfp is the same as l.Panicf but takes a prepared format struct.
func (l *Logger) Panicfp(f *color.Format, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	s := fmt.Sprintf(f.Get(l.color), v...)
	io.WriteString(l.out, s)
	panic(s)
}

// Panic is equivalent to l.Print() followed by a call to panic().
func (l *Logger) Panic(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	s := fmt.Sprint(v...)
	io.WriteString(l.out, s)
	panic(s)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func (l *Logger) Panicln(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	s := fmt.Sprintln(v...)
	io.WriteString(l.out, s)
	panic(s)
}

// SetOutput sets the output destination.
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}

// SetColor sets whether colored output is enabled.
func (l *Logger) SetColor(color bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = color
}

var std = New(os.Stderr, color.IsTerminal(os.Stderr))

// Printf calls the standard Logger's Printf method.
func Printf(format string, v ...interface{}) {
	std.Printf(format, v...)
}

// Print calls the standard Logger's Printf method.
func Print(format string, v ...interface{}) {
	std.Print(v...)
}

// Println calls the standard Logger's Println method.
func Println(format string, v ...interface{}) {
	std.Println(v...)
}

// Fatalf calls the standard Logger's Fatalf method.
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

// Fatal calls the standard Logger's Fatal method.
func Fatal(format string, v ...interface{}) {
	std.Fatal(v...)
}

// Fatalln calls the standard Logger's Fatalln method.
func Fatalln(format string, v ...interface{}) {
	std.Fatalln(v...)
}

// Panicf calls the standard Logger's Panicf method.
func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

// Panic calls the standard Logger's Panic method.
func Panic(format string, v ...interface{}) {
	std.Panic(v...)
}

// Panicln calls the standard Logger's Panicln method.
func Panicln(format string, v ...interface{}) {
	std.Panicln(v...)
}

// Printfp calls the standard Logger's Printfp method.
func Printfp(f *color.Format, v ...interface{}) {
	std.Printfp(f, v...)
}

// Fatalfp calls the standard Logger's Fatalfp method.
func Fatalfp(f *color.Format, v ...interface{}) {
	std.Fatalfp(f, v...)
}

// Panicfp calls the standard Logger's Panicfp method.
func Panicfp(f *color.Format, v ...interface{}) {
	std.Panicfp(f, v...)
}

// SetOutput sets the output destination of the standard Logger.
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// SetColor sets whether colored output is enabled for the standard Logger.
func SetColor(color bool) {
	std.SetColor(color)
}
