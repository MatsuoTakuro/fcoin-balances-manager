package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MatsuoTakuro/fcoin-balances-manager/api/params"
	"github.com/MatsuoTakuro/fcoin-balances-manager/api/shared"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/service"
	"github.com/go-playground/validator/v10"
)

type RegisterUser struct {
	Service   service.RegisterUserService
	Validator *validator.Validate
}

type registerUserReqBody struct {
	Name string `json:"name" validate:"required,min=1,max=20,excludesall='\""`
}

type registerUserRespBody struct {
	UserID  entity.UserID  `json:"user_id"`
	Name    string         `json:"name"`
	Balance shared.Balance `json:"balance"`
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody := &registerUserReqBody{}

	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		err = apperror.DECODE_REQBODY_FAILED.Wrap(err, fmt.Sprintf("failed to decode request body: %q", r.Body))
		apperror.ErrorRespond(ctx, w, err)
		return
	}
	defer r.Body.Close()

	if err := ru.Validator.Struct(reqBody); err != nil {
		err = apperror.BAD_PARAM.Wrap(err,
			fmt.Sprintf("invalid request params: %v", params.InvalidBodyItems(ru.Validator, err)))
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	user, balance, err := ru.Service.RegisterUser(ctx, reqBody.Name)
	if err != nil {
		apperror.ErrorRespond(ctx, w, err)
		return
	}

	respBody := &registerUserRespBody{
		UserID: user.ID,
		Name:   user.Name,
		Balance: shared.Balance{
			Amount: balance.Amount,
		},
	}
	Respond(ctx, w, respBody, http.StatusCreated)
}
