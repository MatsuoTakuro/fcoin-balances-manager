package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
エラーの場合も含めて、この関数で返却するレスポンスを作成する。
*/
func Respond(ctx context.Context, w http.ResponseWriter, body any, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("encode response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		errResp := ErrResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			fmt.Printf("write error-response error: %v", err)
		}
		return
	}

	w.WriteHeader(statusCode)
	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		fmt.Printf("write response error: %v", err)
	}
}
