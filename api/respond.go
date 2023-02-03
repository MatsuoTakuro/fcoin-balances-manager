package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
)

/*
正常時に返却するレスポンスを作成する
*/
func Respond(ctx context.Context, w http.ResponseWriter, respBody any, statusCode int) {

	bodyBytes, err := json.Marshal(respBody)
	if err != nil {
		err = apperror.ENCODE_RESPBODY_FAILED.Wrap(err, "failed to encode response")
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	w.WriteHeader(statusCode)

	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		err = apperror.WRITE_RESPBODY_FAILED.Wrap(err, "failed to write response")
		apperror.ErrorRespond(ctx, w, err)
		return
	}
}
