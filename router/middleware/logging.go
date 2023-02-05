package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/MatsuoTakuro/fcoin-balances-manager/appcontext"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
)

type respLoggingWriter struct {
	w          http.ResponseWriter
	mw         io.Writer
	statusCode int
}

var _ http.ResponseWriter = (*respLoggingWriter)(nil)

func NewRespLoggingWriter(w http.ResponseWriter, buf io.Writer) *respLoggingWriter {
	return &respLoggingWriter{
		w:          w,
		mw:         io.MultiWriter(w, buf),
		statusCode: 0,
	}
}

func (rlw *respLoggingWriter) Header() http.Header {
	return rlw.w.Header()
}

func (rlw *respLoggingWriter) Write(i []byte) (int, error) {
	if rlw.statusCode == 0 {
		rlw.statusCode = http.StatusOK
	}
	return rlw.mw.Write(i)
}

func (rlw *respLoggingWriter) WriteHeader(statusCode int) {
	rlw.statusCode = statusCode
	rlw.w.WriteHeader(statusCode)
}

// 独自のmiddlewareとして、traceID付きのreq/respのloggingを実装
func Logging() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				apperror.ErrorRespond(r.Context(), w,
					apperror.UNKNOWN_ERR.Wrap(err, "faild to read req body for logging"))
			}
			defer r.Body.Close()
			traceID := appcontext.NewTraceID()

			log.Printf("[%d]req: %s %s %s\n", traceID, r.RequestURI, r.Method, reqBody)

			respBodyBuf := &bytes.Buffer{}
			rlw := NewRespLoggingWriter(w, respBodyBuf)
			ctx := appcontext.SetTraceID(r.Context(), traceID)
			r = r.WithContext(ctx)
			r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

			next.ServeHTTP(rlw, r)

			log.Printf("[%d]resp: %d %s\n", traceID, rlw.statusCode, respBodyBuf)
		})
	}
}
