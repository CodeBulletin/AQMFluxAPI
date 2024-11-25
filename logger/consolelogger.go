package logger

import (
	"fmt"
	"sync"
	"time"
)

type consolelogger struct {
	mu     sync.Mutex
}

func NewConsoleLogger() *consolelogger {
	return &consolelogger{
		mu: sync.Mutex{},
	}
}

func (l *consolelogger) Info(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(Cyan)
	fmt.Print("[INFO] ")
	fmt.Print(time.Now().Format("2006/01/02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Println(Reset)
}

func (l *consolelogger) Error(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(Yellow)
	fmt.Print("[ERROR] ")
	fmt.Print(time.Now().Format("2006/01/02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Println(Reset)
}

func (l *consolelogger) Fatal(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(Red)
	fmt.Print("[FATAL] ")
	fmt.Print(time.Now().Format("2006/01/02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Println(Reset)
}

func (l *consolelogger) Debug(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(White)
	fmt.Print("[DEBUG] ")
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Println(Reset)
}

func (l *consolelogger) Status(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(Blue)
	fmt.Print("[STATUS] ")
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Println(Reset)
}

func (l *consolelogger) DBInfo(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(Magenta)
	fmt.Print("[DB INFO] ")
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Print(Reset)
}

func (l *consolelogger) DBError(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(Orange)
	fmt.Print("[DB ERROR] ")
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Println(Reset)
}

func (l *consolelogger) DBFatal(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(Pink)
	fmt.Print("[DB FATAL] ")
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Println(Reset)
}

func (l *consolelogger) DBStatus(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(Green)
	fmt.Print("[DB STATUS] ")
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Println(Reset)
}

func (l *consolelogger) Request(format string, args ... interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Print(Purple)
	fmt.Print("[REQUEST] ")
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print(": ")
	fmt.Printf(format, args...)
	fmt.Println(Reset)
}