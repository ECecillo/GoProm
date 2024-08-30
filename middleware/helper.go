package middleware

import (
	"net/http"

	"github.com/ECecillo/GoProm/types"
)

func CreateStack(xs ...types.Middleware) types.Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}
