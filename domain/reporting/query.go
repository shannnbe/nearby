package reporting

const (
	CreatePrizeQuery = `
		INSERT INTO 
			public.prize (name, period,coupons,session,sequence)
		VALUES('%s', '%s', %v, %v, %v);`

	UpdatePrizeQuery = `
		UPDATE 
			public.prize
		SET 
			"name"='%s', "period"='%s', coupons=%v, "session"=%v, "sequence"=%v
		WHERE id=%v;
	`

	GetCountOfShuffledCouponByPrize = `
		select * from public.total_winner_and_prize
		order by id desc;
	`

	FindPrizeQuery = `SELECT * FROM public.prize WHERE ID = %v`

	GetAllQueryFilterByPeriode = `select * from public.prize where period = '%s' order by sequence ASC`
)
