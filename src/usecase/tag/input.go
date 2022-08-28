package tag

type ListInput struct {
	Keyword string `form:"keyword"`
	Limit   int    `form:"limit,default=10"`
	Status  int    `form:"status"`
}
