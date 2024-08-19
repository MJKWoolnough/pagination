package pagination

const (
	defaultEnd         = 3
	defaultSurrounding = 3
)

// Config is the configuration for a Pagination.
type Config struct {
	Ends, Surrounding uint
}

// New returns a default configuration for Pagination.
func New() Config {
	return Config{
		Ends:        defaultEnd,
		Surrounding: defaultSurrounding,
	}
}

// Get returns the Section information for Pagination.
func (c Config) Get(currPage, lastPage uint) Pagination {
	if currPage > lastPage {
		lastPage = currPage
	}

	p := Pagination{page: currPage}
	start := uint(0)

	for page := uint(0); page <= lastPage; page++ {
		if !(page < c.Ends || // Beginning
			page > lastPage-c.Ends || // End
			((c.Surrounding > currPage || page >= currPage-c.Surrounding) && page <= currPage+c.Surrounding) || // Middle.
			c.Ends > 0 && ((currPage-c.Surrounding-1 == c.Ends && page == c.Ends) || // Merge Beginning and Middle if close enough.
				(currPage+c.Surrounding+1 == lastPage-c.Ends && page == lastPage-c.Ends))) { // Merge Middle and End if close enough.

			if page != start {
				p.sections[p.numSections] = section{start, page - 1}
				p.numSections++
			}

			start = page + 1
		}
	}

	if start < lastPage {
		p.sections[p.numSections] = section{start, lastPage}
		p.numSections++
	}

	return p
}
