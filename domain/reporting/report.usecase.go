package reporting

import (
	"nearby-web-app/common"

	"gofr.dev/pkg/gofr"
)

func UpdatePrizeSequence(ctx *gofr.Context, request PrizeUpdateSequence, Id int) (result interface{}, err error) {
	prize, err := FindPrize(ctx, Id)
	if err != nil {
		return
	}
	if prize.Session < request.Sequence {
		return nil, common.BadRequestError("[PRIZE:UPDATE] sequence larger than session")
	}
	prize.Sequence = request.Sequence
	result, err = UpdatePrize(ctx, prize)
	return
}
