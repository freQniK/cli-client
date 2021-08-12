package middlewares

import (
	"log"
	"net/http"
	"time"

	resttypes "github.com/sentinel-official/cli-client/types/rest"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var (
				rw    = resttypes.NewResponseWriter(w)
				start = time.Now()
			)

			next.ServeHTTP(rw, r)
			log.Printf(
				"- %s - %s %s %s - %d %d - %s - %s - %s",
				r.RemoteAddr,
				r.Proto,
				r.Method,
				r.RequestURI,
				rw.Status,
				rw.Length,
				time.Since(start),
				r.Referer(),
				r.UserAgent(),
			)
		},
	)
}
