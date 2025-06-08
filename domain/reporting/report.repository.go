package reporting

import (
	"fmt"

	"gofr.dev/pkg/gofr"
)

func CreatePrize(ctx *gofr.Context, data PrizeEntity) (result interface{}, err error) {

	prepared_statement := fmt.Sprintf(
		CreatePrizeQuery,
		data.Name,
		data.Period,
		data.Coupons,
		data.Session,
		data.Sequence)

	fmt.Println(prepared_statement)
	return ctx.SQL.ExecContext(ctx, prepared_statement)

}

func UpdatePrize(ctx *gofr.Context, data PrizeEntity) (result interface{}, err error) {
	prepared_statement := fmt.Sprintf(
		UpdatePrizeQuery,
		data.Name,
		data.Period,
		data.Coupons,
		data.Session,
		data.Sequence,
		data.ID)

	fmt.Println(prepared_statement)
	return ctx.SQL.ExecContext(ctx, prepared_statement)
}

func FindPrize(ctx *gofr.Context, ID int) (result PrizeEntity, err error) {

	prepared_statement := fmt.Sprintf(FindPrizeQuery, ID)
	rows, err := ctx.SQL.QueryContext(ctx, prepared_statement)
	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.Period,
			&result.Coupons,
			&result.Session,
			&result.Sequence,
		); err != nil {
			ctx.Logger.Error(err)
			return
		}
	}

	return
}

func GetAllPrize(ctx *gofr.Context, Period string) (results []PrizeEntity, err error) {
	prepared_statement := fmt.Sprintf(GetAllQueryFilterByPeriode, Period)

	rows, err := ctx.SQL.QueryContext(ctx, prepared_statement)
	if err != nil {
		return
	}

	for rows.Next() {
		var result PrizeEntity
		if err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.Period,
			&result.Coupons,
			&result.Session,
			&result.Sequence,
		); err != nil {
			ctx.Logger.Error(err)
			return
		}

		results = append(results, result)
	}

	return
}
