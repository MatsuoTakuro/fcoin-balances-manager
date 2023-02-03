package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MatsuoTakuro/fcoin-balances-manager/api"
	"github.com/MatsuoTakuro/fcoin-balances-manager/api/params"
	"github.com/MatsuoTakuro/fcoin-balances-manager/config"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository/clock"
	"github.com/MatsuoTakuro/fcoin-balances-manager/router/middleware"
	"github.com/MatsuoTakuro/fcoin-balances-manager/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewRouter(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {

	// DBに接続する
	db, cleanup, err := repository.OpenDB(ctx, cfg)
	if err != nil {
		return nil, cleanup, fmt.Errorf("failed to open db: %w", err)
	}

	// その他ルーティングのための各種準備をする
	r := chi.NewRouter()
	v := validator.New()
	c := clock.RealClocker{}
	repo := repository.Repository{Clocker: c}

	// 共通の（独自）ミドルウェアを設定
	r.Use(middleware.RespondingJson())
	r.Use(middleware.Logging())

	// 以下、各API（ハンドラ）に対して、実装されたservice・dbと、公開するpathを割り当てる
	ru := &api.RegisterUser{
		Service: &service.RegisterUserServicer{
			DB:   db,
			Repo: &repo,
		},
		Validator: v,
	}
	r.Post("/user", ru.ServeHTTP)

	gb := &api.GetBalanceDetails{
		Service: &service.GetBalanceDetailsServicer{
			DB:   db,
			Repo: &repo,
		},
		Validator: v,
	}
	ub := &api.UpdateBalance{
		Service: &service.UpdateBalanceServicer{
			DB:   db,
			Repo: &repo,
		},
		Validator: v,
	}
	tc := &api.TransferCoins{
		Service: &service.TransferCoinsServicer{
			DB:   db,
			Repo: &repo,
		},
		Validator: v,
	}
	r.Route("/users", func(sub chi.Router) {
		sub.Get(fmt.Sprintf("/{%s}", params.UserID.Path()), gb.ServeHTTP)
		sub.Patch(fmt.Sprintf("/{%s}", params.UserID.Path()), ub.ServeHTTP)
		sub.Post(fmt.Sprintf("/{%s}/transfer", params.UserID.Path()), tc.ServeHTTP)
	})

	return r, cleanup, nil
}
