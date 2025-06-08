package monitoring

type CouponEntity struct {
}

func (u *CouponEntity) TableName() string {
	return "coupons"
}
