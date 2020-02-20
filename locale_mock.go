// +build unit_test

package locale

import (
	"sync"
)

var mockLang mock

type mock struct {
	s   []string
	err error

	sync.Mutex
}

func (l *mock) get() ([]string, error) {
	l.Lock()
	defer l.Unlock()

	return l.s, l.err
}

func (l *mock) set(s []string, e error) {
	l.Lock()
	defer l.Unlock()
	
	l.s = s
	l.err = e
}

var detectors = []detector{
	mockLang.get,
}
