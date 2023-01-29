package apperror

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/MatsuoTakuro/fcoin-balances-manager/contexts"
)

/*
異常・エラー時に返却するレスポンスを作成する
*/
func ErrorRespond(ctx context.Context, w http.ResponseWriter, err error) {

	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = UnknownErr.Wrap(err, "not found app_error").(*AppError)
	}

	traceID := contexts.GetTracdID(ctx)
	log.Printf("[%d]error: %s: %s -> %s\n",
		traceID, appErr.ErrCode, appErr.ErrMessage, appErr.Err)

	var statusCode int
	switch appErr.ErrCode {
	case NoSelectedData:
		statusCode = http.StatusNotFound
	case DecodeReqBodyFailed, BadParam,
		NoTargetData, RegisterDuplicateDataRestricted:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(appErr)
}
