package api

import (
	"encoding/json"
	"net/http"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/service"
	"github.com/go-playground/validator/v10"
)

type RegisterUser struct {
	Service   service.RegisterUserService
	Validator *validator.Validate
}

type registerUserReqBody struct {
	Name string `json:"name" validate:"required"`
}

type registerUserRespBody struct {
	UserID  entity.UserID   `json:"user_id"`
	Name    string          `json:"name"`
	Balance balanceRespBody `json:"balance"`
}

type balanceRespBody struct {
	Amount uint32 `json:"amount"`
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rb := &registerUserReqBody{}

	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		Respond(ctx, w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	user, balance, err := ru.Service.RegisterUser(ctx, rb.Name)
	if err != nil {
		Respond(ctx, w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
	}

	resp := &registerUserRespBody{
		UserID: user.ID,
		Name:   user.Name,
		Balance: balanceRespBody{
			Amount: balance.Amount,
		},
	}
	Respond(ctx, w, resp, http.StatusOK)
}
