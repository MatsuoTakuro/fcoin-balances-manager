package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MatsuoTakuro/fcoin-balances-manager/api/validation"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/service"
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
	ID          entity.BalanceTransID `json:"id"`
	UserID      entity.UserID         `json:"user_id"`
	BalanceID   entity.BalanceID      `json:"balance_id"`
	Amount      int32                 `json:"amount"`
	ProcessedAt time.Time             `json:"processed_at"`
}

func (ub *UpdateBalance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := validation.UserID.Parse(r)
	if err != nil || userID == 0 {
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	reqBody := &updateBalanceReqBody{}
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		apperror.ErrorRespond(ctx, w,
			apperror.DECODE_REQBODY_FAILED.Wrap(err, fmt.Sprintf("failed to decode request body: %q", r.Body)))
		return
	}
	defer r.Body.Close()

	if err := ub.Validator.Struct(reqBody); err != nil {
		apperror.ErrorRespond(ctx, w,
			apperror.BAD_PARAM.WrapWithErrMessages(err, validation.InvalidItemsErrMessages(ub.Validator, err)))
		return
	}

	balanceTrans, err := ub.Service.UpdateBalance(ctx, userID, reqBody.Amount)
	if err != nil {
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	respBody := &updateBalanceRespBody{
		ID:          balanceTrans.ID,
		UserID:      balanceTrans.UserID,
		BalanceID:   balanceTrans.BalanceID,
		Amount:      balanceTrans.Amount,
		ProcessedAt: balanceTrans.ProcessedAt,
	}
	Respond(ctx, w, respBody, http.StatusOK)
}
