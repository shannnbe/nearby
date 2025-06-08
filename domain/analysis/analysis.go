package analysis

import "time"

type WinnerEntity struct {
	ID              int
	CIF             string
	ShuffleCode     string
	CouponCode      string
	ShuffleDate     time.Time
	Prize           string
	IDPrize         int
	CustomerName    string
	ShuffleSequence int
	ShufflePeriode  int
}

func (u *WinnerEntity) TableName() string {
	return "trx_winner"
}

func (u *WinnerEntity) RestPath() string {
	return "api/v1/winner"
}
