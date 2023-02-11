package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/service"
	handler "github.com/MatsuoTakuro/fcoin-balances-manager/testutil/api"
	"github.com/go-playground/validator/v10"
)

func TestRegisterUser(t *testing.T) {
	t.Parallel()

	type want struct {
		statusCode   int
		respFilePath string
	}

	tests := map[string]struct {
		reqFilePath string
		want        want
	}{
		"ok (created)": {
			reqFilePath: "testdata/register_user/ok_req.json.golden",
			want: want{
				statusCode:   http.StatusCreated,
				respFilePath: "testdata/register_user/ok_resp.json.golden",
			},
		},
		"bad request": {
			reqFilePath: "testdata/register_user/bad_req.json.golden",
			want: want{
				statusCode:   http.StatusBadRequest,
				respFilePath: "testdata/register_user/bad_resp.json.golden",
			},
		},
	}

	for name, sub := range tests {
		sub := sub
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mock := &service.RegisterUserServiceMock{}
			mock.RegisterUserFunc = func(
				ctx context.Context, name string,
			) (*entity.User, *entity.Balance, error) {
				if sub.want.statusCode == http.StatusCreated {
					return &entity.User{
							ID:   1,
							Name: "taro",
						}, &entity.Balance{
							Amount: 0,
						},
						nil
				}
				return nil, nil, errors.New("error from RegisterUserServiceMock")
			}
			ru := RegisterUser{
				Service:   mock,
				Validator: validator.New(),
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/user",
				bytes.NewReader(handler.LoadJSONFile(t, sub.reqFilePath)),
			)

			ru.ServeHTTP(w, r)

			resp := w.Result()
			handler.AssertResponse(t, resp, sub.want.statusCode, handler.LoadJSONFile(t, sub.want.respFilePath))
		})
	}
}
