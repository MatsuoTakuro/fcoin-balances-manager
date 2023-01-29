package middleware

import (
	"log"
	"net/http"

	"github.com/MatsuoTakuro/fcoin-balances-manager/contexts"
)

type respLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewRespLoggingWriter(w http.ResponseWriter) *respLoggingWriter {
	return &respLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rlw *respLoggingWriter) WriteHeader(code int) {
	rlw.code = code
	rlw.ResponseWriter.WriteHeader(code)
}

// 独自のmiddlewareとして、traceID付きのloggingを実装
func Logging() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceID := contexts.NewTraceID()

			log.Printf("[%d]req: %s %s\n", traceID, r.RequestURI, r.Method)

			rlw := NewRespLoggingWriter(w)

			ctx := contexts.SetTraceID(r.Context(), traceID)
			r = r.WithContext(ctx)

			next.ServeHTTP(rlw, r)

			log.Printf("[%d]resp: %d", traceID, rlw.code)
		})
	}
}
