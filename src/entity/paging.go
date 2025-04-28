package entity

type Paging struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
	Limit      int `json:"-"`
	Offset     int `json:"-"`
}

func (p *Paging) ToPaging() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	p.Limit = p.PerPage
	p.Offset = (p.Page - 1) * p.PerPage
}

func (p *Paging) SetTotalPages(counts int64) {
	p.TotalItems = int(counts)
	p.TotalPages = (p.TotalItems + p.PerPage - 1) / p.PerPage
}
