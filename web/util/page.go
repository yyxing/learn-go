package util

type PageParam struct {
	All     bool
	Size    int
	Current int
}

func (p *PageParam) Page() *Page {
	if p.Size <= 0 {
		p.Size = 10
	}
	if p.Size > 24 {
		p.Size = 24
	}
	if p.Current <= 0 {
		p.Current = 1
	}
	return &Page{
		All:     p.All,
		Size:    p.Size,
		Offset:  p.Size * (p.Current - 1),
		Current: p.Current,
	}
}

type Page struct {
	All     bool        `json:"all"`
	Size    int         `json:"size"`
	Offset  int         `json:"offset"`
	Current int         `json:"current"`
	Total   int64       `json:"total"`
	Records interface{} `json:"records"`
}
