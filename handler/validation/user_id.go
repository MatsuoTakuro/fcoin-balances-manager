package validation

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/go-chi/chi/v5"
)

type PathUserID struct {
	Name  string
	regex string
}

var UserID = PathUserID{
	Name:  "user_id",
	regex: "[0-9]+",
}

func (pu PathUserID) Path() string {
	return fmt.Sprintf("%s:%s", pu.Name, pu.regex)
}

func (pu PathUserID) Parse(r *http.Request) (entity.UserID, error) {
	strUserID := chi.URLParam(r, pu.Name)
	userID, err := strconv.ParseInt(strUserID, 10, 64)
	if err != nil {
		return 0, apperror.BAD_PARAM.Wrap(err, fmt.Sprintf("input (%s) in path cannot be parsed into %s", strUserID, pu.Name))
	}
	if userID <= 0 {
		return 0, apperror.NewAppError(apperror.BAD_PARAM, fmt.Sprintf("%s (input: %s) in path cannot be less or 0", pu.Name, strUserID))
	}

	return entity.UserID(userID), nil
}
