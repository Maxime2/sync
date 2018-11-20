// Copyright (c) 2018, Maxim Zakharov. All rights reserved.

package sync

import (
	"runtime"
	"sync"
	"github.com/sirupsen/logrus"
)

type Mutex_trace struct {
	mu     sync.Mutex
	name   string
	logger *logrus.Entry
}

func (m *Mutex_trace) SetNameAndLogger(name string, logger *logrus.Entry) {
	m.name = name
	m.logger = logger
}

func (m *Mutex_trace) Lock() {
	pc, filename, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	m.logger.WithFields(logrus.Fields{
		"name": m.name,
		"status": "request",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("Lock()")
	m.mu.Lock()
	m.logger.WithFields(logrus.Fields{
		"name": m.name,
		"status": "grant",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("Lock()")
}

func (m *Mutex_trace) Unlock() {
	pc, filename, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	m.logger.WithFields(logrus.Fields{
		"name": m.name,
		"status": "request",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("Unlock()")
	m.mu.Unlock()
	m.logger.WithFields(logrus.Fields{
		"name": m.name,
		"status": "grant",
		"file": filename,
		"line" : line,
		"ok": ok,
		"function": details.Name(),
	}).Warn("Unlock()")
}
