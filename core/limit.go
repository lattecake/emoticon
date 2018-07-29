package core

import (
	"net/http"
	"golang.org/x/time/rate"
)

type Limiter struct {
	Lmt *rate.Limiter
}

func NewLimiter(lmt *rate.Limiter) *Limiter {
	return &Limiter{Lmt: lmt}
}

func (l *Limiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if l.Lmt.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (l *Limiter) SetLimit(limit *rate.Limiter) {
	l.Lmt = limit
}
