package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/MatsuoTakuro/fcoin-balances-manager/api/params"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type UpdateBalance struct {
	Service   service.UpdateBalanceService
	Validator *validator.Validate
}

type updateBalanceReqBody struct {
	Amount int32 `json:"amount" validate:"required"`
}

type updateBalanceRespBody struct {
	BalanceTransID entity.BalanceTransID `json:"balance_trans_id"`
	UserID         entity.UserID         `json:"user_id"`
	BalanceID      entity.BalanceID      `json:"balance_id"`
	Amount         int32                 `json:"amount"`
	ProcessedAt    time.Time             `json:"processed_at"`
}

func (ub *UpdateBalance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody := &updateBalanceReqBody{}

	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		err = apperror.DecodeReqBodyFailed.Wrap(err, "failed to decode request body for updating balance")
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	if err := ub.Validator.Struct(reqBody); err != nil {
		err = apperror.BadParam.Wrap(err,
			fmt.Sprintf("invalid request params for updating balance: %v",
				params.InvalidBodyItems(ub.Validator, err)))
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	strUserID := chi.URLParam(r, params.UserID.Name)
	userID, err := strconv.ParseInt(strUserID, 10, 64)
	if err != nil || userID < 1 {
		err = apperror.BadParam.Wrap(err, fmt.Sprintf("invalid request params for updating balance: %s: %s",
			params.UserID, strUserID))
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	balanceTrans, err := ub.Service.UpdateBalance(ctx, entity.UserID(userID), reqBody.Amount)
	if err != nil {
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	respBody := &updateBalanceRespBody{
		BalanceTransID: balanceTrans.ID,
		UserID:         balanceTrans.UserID,
		BalanceID:      balanceTrans.BalanceID,
		Amount:         balanceTrans.Amount,
		ProcessedAt:    balanceTrans.ProcessedAt,
	}
	Respond(ctx, w, respBody, http.StatusCreated)
}
