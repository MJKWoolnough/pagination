package pagination_test

import (
	"fmt"

	"vimagination.zapto.org/pagination"
)

func Example() {
	cfg := pagination.New()
	cfg.Ends = 2
	cfg.Surrounding = 2

	p := cfg.Get(5, 19)

	fmt.Println(p.HTML("/page?"))

	// Output:
	// <a href="/page?1">1</a> <a href="/page?2">2</a> <a href="/page?3">3</a> <a href="/page?4">4</a> <a href="/page?5">5</a> 6 <a href="/page?7">7</a> <a href="/page?8">8</a> ...<a href="/page?19">19</a> <a href="/page?20">20</a>
}
