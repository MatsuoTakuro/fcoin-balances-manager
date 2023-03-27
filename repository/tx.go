package repository

/*
更新系のトランザクション処理についてまとめる
*/

import (
	"context"
	"database/sql"
	"log"

	"github.com/MatsuoTakuro/fcoin-balances-manager/appcontext"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

// トランザクションの共通処理をメソッド化
func (r *Repository) BeginTx(ctx context.Context, db Beginner) (*sql.Tx, func(context.Context, *sql.Tx, error) error, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, apperror.PROCESS_TRANSACTION_FAILED.Wrap(err, err.Error())
	}
	// commitOrRollback関数を返却する
	return tx, commitOrRollback, nil
}

// コミットの実行・コミット失敗時にロールバックする共通処理を関数化
func commitOrRollback(ctx context.Context, tx *sql.Tx, err error) error {
	// トランザクション処理が成功している場合
	if err == nil {
		// コミットを実行
		if err := tx.Commit(); err != nil {
			// コミット失敗の場合、ロールバックを実施
			if rbErr := tx.Rollback(); rbErr != nil {
				// ロールバック失敗の場合、その失敗時のエラーを返却
				return apperror.PROCESS_TRANSACTION_FAILED.Wrap(err, rbErr.Error())
			}
			err = apperror.PROCESS_TRANSACTION_FAILED.Wrap(err, err.Error())
			log.Printf("[%d]trans: rollbacked sucessfully for handling err: %v",
				appcontext.GetTracdID(ctx), err)
			// ロールバック成功の場合、コミット失敗時のエラーを返却
			return err
		}
		// コミット成功の場合、nilを返却
		return nil
		// トランザクション処理が失敗している場合
	} else {
		// ロールバックを実施
		if rbErr := tx.Rollback(); rbErr != nil {
			// ロールバック失敗の場合、その失敗時のエラーを返却
			return apperror.PROCESS_TRANSACTION_FAILED.Wrap(err, rbErr.Error())
		}
		log.Printf("[%d]trans: rollbacked sucessfully for handling err: %v",
			appcontext.GetTracdID(ctx), err)
		// ロールバック成功の場合、トランザクション処理時のエラーを返却
		return err
	}
}

// トランザクション処理の一例
// 新規ユーザの登録と残高作成を行うトランザクション処理
func (r *Repository) RegisterUserTx(
	ctx context.Context, db Beginner, name string,
	// 返り値のerrorをerr変数で定義しておく
) (user *entity.User, balance *entity.Balance, err error) {

	// トランザクションの共通処理を呼び出し、commitOrRollback関数を受け取る
	tx, commitOrRollback, err := r.BeginTx(ctx, db)
	defer func() {
		// defer関数内でcommitOrRollback関数を呼び出す
		// errとは別に、txErr変数を定義・初期化する
		if txErr := commitOrRollback(ctx, tx, err); txErr != nil {
			// 返り値であるerrにtxErrを代入（上書き）する
			// 最終的に、このerrが本メソッド（RegisterUserTx）の呼び出し元に返却される
			err = txErr
		}
	}()

	user, err = r.CreateUser(ctx, tx, name)
	if err != nil {
		// defer関数へ（errを上書く可能性あり）
		return nil, nil, err
	}

	balance, err = r.CreateBalance(ctx, tx, user.ID)
	if err != nil {
		// defer関数へ（errを上書く可能性あり）
		return nil, nil, err
	}
	// defer関数へ（nilを上書く可能性あり）
	return user, balance, nil
}

// （コイン転送によらない）残高更新トランザクション処理
func (r *Repository) UpdateBalanceTx(
	ctx context.Context, db Beginner, balance *entity.Balance, amount int32,
) (balanceTrans *entity.BalanceTrans, err error) {

	tx, commitOrRollback, err := r.BeginTx(ctx, db)
	defer func() {
		if txErr := commitOrRollback(ctx, tx, err); txErr != nil {
			err = txErr
		}
	}()

	balanceTrans, err = r.CreateBalanceTrans(ctx, tx, balance.UserID, balance.ID, amount)
	if err != nil {
		return nil, err
	}

	err = r.UpdateBalanceByID(ctx, tx, balance.ID, balance.Amount)
	if err != nil {
		return nil, err
	}

	return balanceTrans, nil
}

func (r *Repository) TransferCoinsTx(
	ctx context.Context, db Beginner, fromBalance *entity.Balance, toBalance *entity.Balance, amount uint32,
) (*entity.BalanceTrans, error) {

	tx, commitOrRollback, err := r.BeginTx(ctx, db)
	defer func() {
		if txErr := commitOrRollback(ctx, tx, err); txErr != nil {
			err = txErr
		}
	}()

	transferTrans, err := r.CreateTransferTrans(
		ctx, tx, fromBalance, toBalance, amount)
	if err != nil {
		return nil, err
	}

	balanceTrans, err := r.CreateBalanceTransByTransfer(
		ctx, tx, transferTrans.FromUser, transferTrans.FromBalance, transferTrans.ID, -int32(amount))
	if err != nil {
		return nil, err
	}

	err = r.UpdateBalanceByID(
		ctx, tx, fromBalance.ID, fromBalance.Amount)
	if err != nil {
		return nil, err
	}

	_, err = r.CreateBalanceTransByTransfer(
		ctx, tx, transferTrans.ToUser, transferTrans.ToBalance, transferTrans.ID, int32(amount))
	if err != nil {
		return nil, err
	}

	err = r.UpdateBalanceByID(
		ctx, tx, toBalance.ID, toBalance.Amount)
	if err != nil {
		return nil, err
	}

	balanceTrans.TransferTrans = *transferTrans

	return balanceTrans, nil
}
