// Package pagination implements a pagination solution for multiple front-ends
package pagination

const elipses = "..."

// Pagination contains the information necessary to print a proper pagination
type Pagination struct {
	numSections byte
	page        uint
	sections    [3]section
}

// Print converts the Pagination sections into a string.
//
// The pageFn func takes a page number and returns whatever text is needed for
// that page.
//
// The sep string is what is to appear between the sections.
func (p Pagination) Print(pageFn func(uint) string, sep string) string {
	str := ""
	for i := byte(0); i < p.numSections; i++ {
		if i != 0 {
			str += sep
		}
		for page := p.sections[i].First; page <= p.sections[i].Last; page++ {
			str += pageFn(page)
		}
	}
	return str
}

// HTML calls Print with a HTML based pageFn and a simple elipses.
//
// The urlBase will have the page number appended to it, so it needs to be
// formatted with this in mind.
func (p Pagination) HTML(urlBase string) string {
	return p.Print(func(page uint) string {
		numStr := itoa(page + 1)
		if page == p.page {
			return numStr + " "
		}
		return "<a href=\"" + urlBase + numStr + "\">" + numStr + "</a> "
	}, elipses)
}

// String stringifies the Sections with a simple pageFn
func (p Pagination) String() string {
	return p.Print(func(page uint) string {
		return itoa(page+1) + " "
	}, elipses)
}

// Section contains the information for a single section of a pagination.
type section struct {
	First, Last uint
}

func itoa(n uint) string {
	if n == 0 {
		return "0"
	}
	var digits [20]byte
	pos := 20
	for ; n > 0; n /= 10 {
		pos--
		digits[pos] = '0' + byte(n%10)
	}
	return string(digits[pos:])
}
