package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MatsuoTakuro/fcoin-balances-manager/api"
	"github.com/MatsuoTakuro/fcoin-balances-manager/config"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository/clock"
	"github.com/MatsuoTakuro/fcoin-balances-manager/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewRouter(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {

	// DBに接続する。
	db, cleanup, err := repository.OpenDB(ctx, cfg)
	if err != nil {
		return nil, cleanup, fmt.Errorf("failed to open db: %w", err)
	}

	// その他ルーティングのための各種準備をする。
	r := chi.NewRouter()
	v := validator.New()
	c := clock.RealClocker{}
	repo := repository.Repository{Clocker: c}

	// 以下、各API（ハンドラ）に対して、公開するパスを割り当てる。
	ru := &api.RegisterUser{
		Service: &service.RegisterUserServiceImpl{
			DB:   db,
			Repo: &repo,
		},
		Validator: v,
	}
	r.Post("/user", ru.ServeHTTP)

	return r, cleanup, nil
}
