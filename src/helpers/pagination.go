package helpers

type Paging struct {
	Page       int32 `json:"page"`
	PerPage    int32 `json:"per_page"`
	TotalPages int32 `json:"total_pages"`
	TotalItems int32 `json:"total_items"`
}

func ToPaging(page, perPage string) Paging {
	paging := Paging{
		Page:       1,
		PerPage:    5,
		TotalPages: 0,
		TotalItems: 0,
	}
	pageInt, err := StringToInt32(page)
	if err == nil {
		paging.Page = pageInt
	}

	perPageInt, err := StringToInt32(perPage)
	if err == nil {
		paging.PerPage = perPageInt
	}
	return paging
}
