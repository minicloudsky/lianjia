package pagination

type Pagination struct {
	Page     int
	PageSize int
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}
