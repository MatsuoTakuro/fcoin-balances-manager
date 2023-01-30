package apperror

/*
クライアント側に返却するエラーコードを独自に定義する。
api、service、repositoryの各層で発生したエラーもここで管理する（各層で独自errorを定義する方法も考えられる）。
*/
type ErrCode string

const (
	UnknownErr ErrCode = "unknown_error"

	/*
		api層のエラー
	*/

	DecodeReqBodyFailed  ErrCode = "decode_req_body_failed"  // リクエストデータの読み込みに失敗
	BadParam             ErrCode = "bad_param"               // リクエストデータが不適切
	EncodeRespBodyFailed ErrCode = "encode_resp_body_failed" // レスポンスデータの変換(JSON)に失敗
	WriteRespBodyFailed  ErrCode = "write_resp_body_failed"  // レスポンスデータの書き込みに失敗

	/*
		service層のエラー
	*/

	ConsumedAmountOverBalance ErrCode = "consumed_amount_over_balance" // 残高を超える消費・使用
	OverMaxBalanceLimit       ErrCode = "over_max_balance_limit"       // 追加される残高の合計が上限額を超える

	/*
		repository層のエラー
	*/

	RegisterDataFailed              ErrCode = "register_data_failed"               // データ登録自体に失敗
	RegisterDuplicateDataRestricted ErrCode = "register_duplicate_data_restricted" // 重複データであるため登録不可
	GetDataFailed                   ErrCode = "get_data_failed"                    // データ取得自体に失敗
	NoSelectedData                  ErrCode = "no_selected_data"                   // データ取得を実行したが0件だった
	UpdateDataFailed                ErrCode = "update_data_failed"                 // データ更新自体に失敗
	NoTargetData                    ErrCode = "no_target_data"                     // 更新したい対象のデータが0件だった
)

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{
		ErrCode:    code,
		ErrMessage: message,
		Err:        err,
	}
}
