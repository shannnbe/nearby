package analysis

import "time"

type WinnerResponse struct {
	ID              int       `json:"id" sql:"auto_increment"`
	CIF             string    `json:"cif"`
	ShuffleCode     string    `json:"shuffle_code"`
	CouponCode      string    `json:"coupon_code"`
	ShuffleDate     time.Time `json:"shuffle_date"`
	Prize           string    `json:"prize"`
	IDPrize         int       `json:"id_prize"`
	CustomerName    string    `json:"customer_name"`
	ShuffleSequence int       `json:"shuffle_sequence"`
	ShufflePeriode  int       `json:"shuffle_periode"`
}
