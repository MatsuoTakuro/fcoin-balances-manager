package apperror

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/MatsuoTakuro/fcoin-balances-manager/appcontext"
)

/*
異常・エラー時に返却するレスポンスを作成する
*/
func ErrorRespond(ctx context.Context, w http.ResponseWriter, err error) {

	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = UNKNOWN_ERR.Wrap(err, "not found app_error").(*AppError)
	}

	traceID := appcontext.GetTracdID(ctx)
	log.Printf("[%d]error: %s: %s -> %s\n",
		traceID, appErr.ErrCode, appErr.ErrMessage, appErr.Err)

	var statusCode int
	switch appErr.ErrCode {
	case NO_SELECTED_DATA:
		statusCode = http.StatusNotFound
	case DECODE_REQBODY_FAILED, BAD_PARAM,
		CONSUMED_AMOUNT_OVER_BALANCE, OVER_MAX_BALANCE_LIMIT,
		NO_TARGET_DATA, REGISTER_DUPLICATE_DATA_RESTRICTED:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(appErr)
}
