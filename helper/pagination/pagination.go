package pagination

type Pagination struct {
	Limit        int         	`json:"limit,omitempty;query:limit"`
	Page         int         	`json:"page,omitempty;query:page"`
	Sort         string      	`json:"sort,omitempty;query:sort"`
	TotalRows    int64       	`json:"total_rows"`
	TotalPages   int         	`json:"total_pages"`
	NextPage   	 bool         	`json:"next_page"`
	PrevPage   	 bool         	`json:"prev_page"`
	Contents     interface{}	`json:"contents"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
