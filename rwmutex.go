// Copyright (c) 2018, Maxim Zakharov. All rights reserved.

package sync

import (
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
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "request",
	}).Error("RLock()")
	rw.mu.RLock()
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "grant",
	}).Error("RLock()")
}

func (rw *RWMutex_trace) RUnlock() {
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "request",
	}).Error("RUnlock()")
	rw.mu.RUnlock()
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "grant",
	}).Error("RUnlock()")
}

func (rw *RWMutex_trace) Lock() {
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "request",
	}).Error("Lock()")
	rw.mu.Lock()
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "grant",
	}).Error("Lock()")
}

func (rw *RWMutex_trace) Unlock() {
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "request",
	}).Error("Unlock()")
	rw.mu.Unlock()
	rw.logger.WithFields(logrus.Fields{
		"name": rw.name,
		"status": "grant",
	}).Error("Unlock()")
}
