package monitoring

type CouponRequest struct {
	Year   int    `json:"year"`
	Quarter string `json:"quarter"`
}

type CouponRequestWithDetailRequest struct {
	PrizeId int `json:"prize_id"`
	Coupons int `json:"coupons"`
}

type PrizeResponse struct {
	ID      int
	Coupons int
}

type CouponResponse struct {
	CouponCode    string `json:"coupon_code"`
	CIF           string `json:"cif"`
	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
	ProductName   string `json:"product_name"`
	Period        string `json:"period"`
	Prize         string `json:"prize"`
	MaskedName    string `json:"name"`
}

type CouponShuffleResponse struct {
	Sequence      int    `json:"no_urut"`
	CIF           string `json:"cif"`
	CouponCode    string `json:"kupon_undian"`
	CustomerName  string `json:"nama_nasabah"`
	PrizeSequence int    `json:"urut_undi"`
}

type CouponRandomizeRequest struct {
	Quarter  string `json:"quarter"`
	Coupons int    `json:"coupons"`
}

type CouponRandomize struct {
	Sequence     int    `json:"no_urut"`
	CIF          string `json:"cif"`
	CouponCode   string `json:"kupon_undian"`
	CustomerName string `json:"nama_nasabah"`
}

type CountShuffledCouponResponse struct {
	IDPrize     int    `json:"id_prize"`
	PrizeName   string `json:"prize_name"`
	TotalWinner int    `json:"total_winner"`
	TotalPrize  int    `json:"total_prize"`
}
