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
	DecodeReqBodyFailed  ErrCode = "decode_req_body_failed"
	BadParam             ErrCode = "bad_param" // リクエストデータが不適切
	EncodeRespBodyFailed ErrCode = "encode_resp_body_failed"
	WriteRespBodyFailed  ErrCode = "write_resp_body_failed"

	/*
		service層のエラー
	*/
	AmountOverBalance ErrCode = "amount_over_balance"

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
