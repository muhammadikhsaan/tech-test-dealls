package modules

import (
	"net/http"
	"strconv"
	"time"

	"dealls.test/material/src/core"
	"dealls.test/material/src/static"
)

type ThrottleValue map[string]time.Time

type throttle struct {
	data ThrottleValue
}

type Throttle interface {
	Get(key string) *time.Time
	Insert(key string) *core.Error
	Delete(key string)
}

func NewThrottle() Throttle {
	return &throttle{
		data: ThrottleValue{},
	}
}

func (t *throttle) Get(key string) *time.Time {
	if v, ok := t.data[key]; ok {
		return &v
	}

	return nil
}

func (t *throttle) Insert(key string) *core.Error {
	timeout, err := strconv.Atoi(static.REQUEST_TIMEOUT)

	if err != nil {
		return &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to convert timeout value",
		}
	}

	t.data[key] = time.Now().Add(time.Duration(timeout) * time.Second)
	return nil
}

func (t *throttle) Delete(key string) {
	delete(t.data, key)
}
