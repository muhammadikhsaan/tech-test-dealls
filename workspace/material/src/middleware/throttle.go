package middleware

import (
	"fmt"
	"net/http"
	"sync"

	"dealls.test/material/src/contract"
	"dealls.test/material/src/core"
	"dealls.test/material/src/modules"
	"dealls.test/material/src/static"
)

var (
	once sync.Once
	t    modules.Throttle
)

func init() {
	once.Do(func() {
		t = modules.NewThrottle()
	})
}

func Throttle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := core.New(w, r)
		identity := ctx.Header().Get(static.HEADER_REQUEST_IDENTITY_KEY)

		if identity == "" {
			ctx.JSON(http.StatusBadRequest, &contract.ResponseError{
				Message: fmt.Sprintf("Header %s is required", static.HEADER_REQUEST_IDENTITY_KEY),
			})
			return
		}

		if exist := t.Get(identity); exist != nil {
			ctx.JSON(http.StatusBadRequest, &contract.ResponseError{
				Message: "Too many requests for the same",
			})
			return
		}

		t.Insert(identity)
		defer t.Delete(identity)

		h.ServeHTTP(w, r)
	})
}
