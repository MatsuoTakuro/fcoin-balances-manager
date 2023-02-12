package apperror

/*
クライアント側に返却するエラーコードを独自に定義する。
api、service、repositoryの各層で発生したエラーもここで管理する（各層で独自errorを定義する方法も考えられる）。
*/
type ErrCode string

const (
	UNKNOWN_ERR ErrCode = "unknown_error"

	/*
		handler層のエラー
	*/

	DECODE_REQBODY_FAILED  ErrCode = "decode_req_body_failed"  // リクエストデータの読み込みに失敗
	BAD_PARAM              ErrCode = "bad_param"               // リクエストデータが不適切
	ENCODE_RESPBODY_FAILED ErrCode = "encode_resp_body_failed" // レスポンスデータの変換(JSON)に失敗
	WRITE_RESPBODY_FAILED  ErrCode = "write_resp_body_failed"  // レスポンスデータの書き込みに失敗

	/*
		service層のエラー
	*/

	CONSUMED_AMOUNT_OVER_BALANCE       ErrCode = "consumed_amount_over_balance"      // 残高を超える消費・使用
	OVER_MAX_BALANCE_LIMIT             ErrCode = "over_max_balance_limit"            // 追加される残高の合計が上限額を超える
	NO_TRANSFERRING_COINS_BY_SAME_USER ErrCode = "no_transfer_of_coins_by_same_user" // 同一ユーザーによるコイン転送の禁止

	/*
		repository層のエラー
	*/

	REGISTER_DATA_FAILED               ErrCode = "register_data_failed"               // データ登録自体に失敗
	REGISTER_DUPLICATE_DATA_RESTRICTED ErrCode = "register_duplicate_data_restricted" // 重複データであるため登録不可
	GET_DATA_FAILED                    ErrCode = "get_data_failed"                    // データ取得自体に失敗
	NO_SELECTED_DATA                   ErrCode = "no_selected_data"                   // データ取得を実行したが0件だった
	UPDATE_DATA_FAILED                 ErrCode = "update_data_failed"                 // データ更新自体に失敗
	NO_TARGET_DATA                     ErrCode = "no_target_data"                     // 更新したい対象のデータが0件だった
	PROCESS_TRANSACTION_FAILED         ErrCode = "process_transaction_failed"         // トランザクション処理に失敗
)

func (code ErrCode) Wrap(err error, messages ...string) error {
	return &AppError{
		ErrCode:     code,
		ErrMessages: messages,
		Err:         err,
	}
}

func (code ErrCode) WrapWithErrMessages(err error, messages []string) error {
	return &AppError{
		ErrCode:     code,
		ErrMessages: messages,
		Err:         err,
	}
}
