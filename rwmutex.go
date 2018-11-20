// Copyright (c) 2018, Maxim Zakharov. All rights reserved.

package sync

import (
	"runtime"
	"sync"
	"github.com/sirupsen/logrus"
)

type RWMutex_trace struct {
	mu     sync.RWMutex
	name   string
	logger *logrus.Entry
}

func (rw *RWMutex_trace) SetNameAndLogger(name string, logger *logrus.Entry) {
	rw.name = name
	rw.logger = logger
}

func (rw *RWMutex_trace) RLock() {
	pc, filename, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "request",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("RLock()")
	rw.mu.RLock()
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "grant",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("RLock()")
}

func (rw *RWMutex_trace) RUnlock() {
	pc, filename, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "request",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("RUnlock()")
	rw.mu.RUnlock()
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "grant",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("RUnlock()")
}

func (rw *RWMutex_trace) Lock() {
	pc, filename, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "request",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("Lock()")
	rw.mu.Lock()
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "grant",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("Lock()")
}

func (rw *RWMutex_trace) Unlock() {
	pc, filename, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "request",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("Unlock()")
	rw.mu.Unlock()
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "grant",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("Unlock()")
}
