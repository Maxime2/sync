// Copyright (c) 2018, Maxim Zakharov. All rights reserved.

package sync

import (
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
	m.logger.WithFields(logrus.Fields{
		"name": m.name,
		"status": "request",
	}).Warn("Lock()")
	m.mu.Lock()
	m.logger.WithFields(logrus.Fields{
		"name": m.name,
		"status": "grant",
	}).Warn("Lock()")
}

func (m *Mutex_trace) Unlock() {
	m.logger.WithFields(logrus.Fields{
		"name": m.name,
		"status": "request",
	}).Warn("Unlock()")
	m.mu.Unlock()
	m.logger.WithFields(logrus.Fields{
		"name": m.name,
		"status": "grant",
	}).Warn("Unlock()")
}
