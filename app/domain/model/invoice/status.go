package invoice

// status : 請求書のステータスの列挙型を定義
type status int

const (
	// statusPending : 0 未処理
	statusPending status = iota
	// statusProcessing : 1 処理中
	statusProcessing
	// statusPaid : 2: 支払い済み
	statusPaid
	// statusError : 3: エラー
	statusError
)

func (s status) String() string {
	switch s {
	case statusPending:
		return "未処理"
	case statusProcessing:
		return "処理中"
	case statusPaid:
		return "支払い済み"
	case statusError:
		return "エラー"
	}

	return ""
}
