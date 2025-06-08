package reporting

type PrizeEntity struct {
	ID       int    `json:"id"  sql:"auto_increment"`
	Name     string `json:"name"`
	Period   string `json:"period"`
	Coupons  int    `json:"coupons"`
	Session  int    `json:"session"`
	Sequence int    `json:"sequence"`
}

func (u *PrizeEntity) TableName() string {
	return "prize"
}

func (u *PrizeEntity) RestPath() string {
	return "api/v1/prizes"
}
