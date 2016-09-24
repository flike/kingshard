// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package golog

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//log level, from low to high, more high means more serious
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	Ltime  = 1 << iota //time format "2006/01/02 15:04:05"
	Lfile              //file.go:123
	Llevel             //[Trace|Debug|Info...]
)

var LevelName [6]string = [6]string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

const (
	LogSqlOn       = "on"
	LogSqlOff      = "off"
	TimeFormat     = "2006/01/02 15:04:05"
	maxBufPoolSize = 16
)

type Logger struct {
	sync.Mutex

	level int
	flag  int

	handler Handler

	quit chan struct{}
	msg  chan []byte

	bufs [][]byte

	wg sync.WaitGroup

	closed bool
}

//new a logger with specified handler and flag
func New(handler Handler, flag int) *Logger {
	var l = new(Logger)

	l.level = LevelInfo
	l.handler = handler

	l.flag = flag

	l.quit = make(chan struct{})
	l.closed = false

	l.msg = make(chan []byte, 1024)

	l.bufs = make([][]byte, 0, 16)

	l.wg.Add(1)
	go l.run()

	return l
}

//new a default logger with specified handler and flag: Ltime|Lfile|Llevel
func NewDefault(handler Handler) *Logger {
	return New(handler, Ltime|Lfile|Llevel)
}

func newStdHandler() *StreamHandler {
	h, _ := NewStreamHandler(os.Stdout)
	return h
}

var std = NewDefault(newStdHandler())

func Close() {
	std.Close()
}

func (l *Logger) run() {
	defer l.wg.Done()
	for {
		select {
		case msg := <-l.msg:
			l.handler.Write(msg)
			l.putBuf(msg)
		case <-l.quit:
			if len(l.msg) == 0 {
				return
			}
		}
	}
}

func (l *Logger) popBuf() []byte {
	l.Lock()
	var buf []byte
	if len(l.bufs) == 0 {
		buf = make([]byte, 0, 1024)
	} else {
		buf = l.bufs[len(l.bufs)-1]
		l.bufs = l.bufs[0 : len(l.bufs)-1]
	}
	l.Unlock()

	return buf
}

func (l *Logger) putBuf(buf []byte) {
	l.Lock()
	if len(l.bufs) < maxBufPoolSize {
		buf = buf[0:0]
		l.bufs = append(l.bufs, buf)
	}
	l.Unlock()
}

func (l *Logger) Close() {
	if l.closed {
		return
	}
	l.closed = true

	close(l.quit)
	l.wg.Wait()
	l.quit = nil

	l.handler.Close()
}

//set log level, any log level less than it will not log
func (l *Logger) SetLevel(level int) {
	l.level = level
}

func (l *Logger) Level() int {
	return l.level
}

//a low interface, maybe you can use it for your special log format
//but it may be not exported later......
func (l *Logger) Output(callDepth int, level int, format string, v ...interface{}) {
	if l.level > level {
		return
	}

	buf := l.popBuf()

	if l.flag&Ltime > 0 {
		now := time.Now().Format(TimeFormat)
		buf = append(buf, now...)
		buf = append(buf, " - "...)
	}

	if l.flag&Llevel > 0 {
		buf = append(buf, LevelName[level]...)
		buf = append(buf, " - "...)
	}

	if l.flag&Lfile > 0 {
		_, file, line, ok := runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		} else {
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					file = file[i+1:]
					break
				}
			}
		}

		buf = append(buf, file...)
		buf = append(buf, ":["...)

		buf = strconv.AppendInt(buf, int64(line), 10)
		buf = append(buf, "] - "...)
	}

	s := fmt.Sprintf(format, v...)

	buf = append(buf, s...)

	if s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}

	l.msg <- buf
}

func SetLevel(level int) {
	std.SetLevel(level)
}

func StdLogger() *Logger {
	return std
}

func GetLevel() int {
	return std.level
}

//全局变量
var GlobalSysLogger *Logger = StdLogger()
var GlobalSqlLogger *Logger = GlobalSysLogger

func (l *Logger) Write(p []byte) (n int, err error) {
	output(LevelInfo, "web", "api", string(p), 0)
	return len(p), nil
}

func escape(s string, filterEqual bool) string {
	dest := make([]byte, 0, 2*len(s))
	for i := 0; i < len(s); i++ {
		r := s[i]
		switch r {
		case '|':
			continue
		case '%':
			dest = append(dest, '%', '%')
		case '=':
			if !filterEqual {
				dest = append(dest, '=')
			}
		default:
			dest = append(dest, r)
		}
	}

	return string(dest)
}

func OutputSql(state string, format string, v ...interface{}) {
	l := GlobalSqlLogger
	buf := l.popBuf()

	if l.flag&Ltime > 0 {
		now := time.Now().Format(TimeFormat)
		buf = append(buf, now...)
		buf = append(buf, " - "...)
	}

	if l.flag&Llevel > 0 {
		buf = append(buf, state...)
		buf = append(buf, " - "...)
	}

	s := fmt.Sprintf(format, v...)

	buf = append(buf, s...)

	if s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}

	l.msg <- buf
}

func output(level int, module string, method string, msg string, reqId uint32, args ...interface{}) {
	if level < GlobalSysLogger.Level() {
		return
	}

	num := len(args) / 2
	var argsBuff bytes.Buffer
	for i := 0; i < num; i++ {
		argsBuff.WriteString(escape(fmt.Sprintf("%v=%v", args[i*2], args[i*2+1]), false))
		if (i+1)*2 != len(args) {
			argsBuff.WriteString("|")
		}
	}
	if len(args)%2 == 1 {
		argsBuff.WriteString(escape(fmt.Sprintf("%v", args[len(args)-1]), false))
	}

	content := fmt.Sprintf(`[%s] "%s" "%s" "%s" conn_id=%d`,
		module, method, msg, argsBuff.String(), reqId)

	GlobalSysLogger.Output(3, level, content)
}

func Trace(module string, method string, msg string, reqId uint32, args ...interface{}) {
	output(LevelTrace, module, method, msg, reqId, args...)
}
func Debug(module string, method string, msg string, reqId uint32, args ...interface{}) {
	output(LevelDebug, module, method, msg, reqId, args...)
}
func Info(module string, method string, msg string, reqId uint32, args ...interface{}) {
	output(LevelInfo, module, method, msg, reqId, args...)
}
func Warn(module string, method string, msg string, reqId uint32, args ...interface{}) {
	output(LevelWarn, module, method, msg, reqId, args...)
}
func Error(module string, method string, msg string, reqId uint32, args ...interface{}) {
	output(LevelError, module, method, msg, reqId, args...)
}
func Fatal(module string, method string, msg string, reqId uint32, args ...interface{}) {
	output(LevelFatal, module, method, msg, reqId, args...)
}
