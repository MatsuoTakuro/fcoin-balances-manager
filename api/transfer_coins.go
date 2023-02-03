package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/MatsuoTakuro/fcoin-balances-manager/api/params"
	"github.com/MatsuoTakuro/fcoin-balances-manager/api/shared"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type TransferCoins struct {
	Service   service.TransferCoinsService
	Validator *validator.Validate
}

type transferCoinsReqBody struct {
	UserID entity.UserID `json:"user_id" validate:"required"`
	Amount uint32        `json:"amount" validate:"required,min=1"`
}

type transferCoinsRespBody struct {
	ID          entity.BalanceTransID `json:"id"`
	UserID      entity.UserID         `json:"user_id"`
	BalanceID   entity.BalanceID      `json:"balance_id"`
	Transfer    shared.TransferTrans  `json:"transfer"`
	Amount      int32                 `json:"amount"`
	ProcessedAt time.Time             `json:"processed_at"`
}

func (tc *TransferCoins) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody := &transferCoinsReqBody{}

	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		err = apperror.DECODE_REQBODY_FAILED.Wrap(err, "failed to decode request body for trasferring coins")
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	if err := tc.Validator.Struct(reqBody); err != nil {
		err = apperror.BAD_PARAM.Wrap(err,
			fmt.Sprintf("invalid request params for trasferring coins: %v",
				params.InvalidBodyItems(tc.Validator, err)))
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	strUserID := chi.URLParam(r, params.UserID.Name)
	userID, err := strconv.ParseInt(strUserID, 10, 64)
	if err != nil || userID < 1 {
		err = apperror.BAD_PARAM.Wrap(err, fmt.Sprintf("invalid request params for trasferring coins: %s: %s",
			params.UserID, strUserID))
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	balanceTrans, err := tc.Service.TransferCoins(ctx, entity.UserID(userID), reqBody.UserID, reqBody.Amount)
	if err != nil {
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	respBody := &transferCoinsRespBody{
		ID:        balanceTrans.ID,
		UserID:    balanceTrans.UserID,
		BalanceID: balanceTrans.BalanceID,
		Transfer: shared.TransferTrans{
			ID:          balanceTrans.TransferTrans.ID,
			FromUser:    balanceTrans.TransferTrans.FromUser,
			FromBalance: balanceTrans.TransferTrans.FromBalance,
			ToUser:      balanceTrans.TransferTrans.ToUser,
			ToBalance:   balanceTrans.TransferTrans.ToBalance,
			Amount:      balanceTrans.TransferTrans.Amount,
			ProcessedAt: balanceTrans.TransferTrans.ProcessedAt,
		},
		Amount:      balanceTrans.Amount,
		ProcessedAt: balanceTrans.ProcessedAt,
	}
	Respond(ctx, w, respBody, http.StatusOK)
}
