package entity

type Paging struct {
	Page       int32 `json:"page"`
	PerPage    int32 `json:"per_page"`
	TotalPages int32 `json:"total_pages"`
	TotalItems int32 `json:"total_items"`
	Limit      int32 `json:"-"`
	Offset     int32 `json:"-"`
}

func ToPaging(page, perPage int32) Paging {
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	return Paging{
		Page:    page,
		PerPage: perPage,
		Limit:   perPage,
		Offset:  (page - 1) * perPage,
	}
}

func (p *Paging) SetTotalPages(counts int64) {
	p.TotalItems = int32(counts)
	p.TotalPages = (p.TotalItems + p.PerPage - 1) / p.PerPage
}
