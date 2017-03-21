package stackexchange

// Paging - https://api.stackexchange.com/docs/paging
type Paging struct {
	page       int
	pageSize   int
	hasMore    bool
	curentPage int
}

// SetPage sets current page
func (p *Paging) SetPage(page int) {
	p.page = page
}

// GetCurrentPageNr return current page number
func (p *Paging) GetCurrentPageNr() int {
	if p.page == 0 {
		p.page = 1
	}
	return p.page
}

// SetPageSize sets current page
func (p *Paging) SetPageSize(pageSize int) {
	p.pageSize = pageSize
}

// NextPage increases page number
func (p *Paging) NextPage() {
	p.page++
}

// PreviousPage decreases page number
func (p *Paging) PreviousPage() {
	if p.page > 0 {
		p.page--
	}
}

// FirstPage sets current page to 1
func (p *Paging) FirstPage() {
	p.page = 1
}

// HasMore either there are more results
func (p *Paging) HasMore() bool {
	return p.hasMore
}
