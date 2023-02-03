package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MatsuoTakuro/fcoin-balances-manager/api/params"
	"github.com/MatsuoTakuro/fcoin-balances-manager/api/shared"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type GetBalanceDetails struct {
	Service   service.GetBalanceDetails
	Validator *validator.Validate
}

type getBalanceDetailsRespBody struct {
	Balance shared.Balance        `json:"balance"`
	History []shared.BalanceTrans `json:"history"`
}

func (gb *GetBalanceDetails) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	strUserID := chi.URLParam(r, params.UserID.Name)
	userID, err := strconv.ParseInt(strUserID, 10, 64)
	if err != nil || userID < 1 {
		err = apperror.BAD_PARAM.Wrap(err, fmt.Sprintf("invalid request params for getting balance details: %s: %s",
			params.UserID, strUserID))
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	balance, history, err := gb.Service.GetBalanceDetails(ctx, entity.UserID(userID))
	if err != nil {
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	respBody := &getBalanceDetailsRespBody{
		Balance: shared.Balance{
			Amount: balance.Amount,
		},
	}

	for _, bt := range history {
		balanceTrans := &shared.BalanceTrans{
			ID:          bt.ID,
			UserID:      bt.UserID,
			BalanceID:   bt.BalanceID,
			Amount:      bt.Amount,
			ProcessedAt: bt.ProcessedAt,
		}

		if bt.TransferTrans.ID == 0 {
			respBody.History = append(respBody.History, *balanceTrans)
			continue
		}

		balanceTrans.Transfer = &shared.TransferTrans{
			ID:          bt.TransferTrans.ID,
			FromUser:    bt.TransferTrans.FromUser,
			FromBalance: bt.TransferTrans.FromBalance,
			ToUser:      bt.TransferTrans.ToUser,
			ToBalance:   bt.TransferTrans.ToBalance,
			Amount:      bt.TransferTrans.Amount,
			ProcessedAt: bt.TransferTrans.ProcessedAt,
		}
		respBody.History = append(respBody.History, *balanceTrans)
	}

	Respond(ctx, w, respBody, http.StatusOK)
}
