# pagination

[![CI](https://github.com/MJKWoolnough/pagination/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/pagination/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/pagination.svg)](https://pkg.go.dev/vimagination.zapto.org/pagination)

--
    import "vimagination.zapto.org/pagination"

Package pagination implements a pagination solution for multiple front-ends.

## Highlights

 - Configure how many page numbers surround the selected page and how many are listed at the beginning and end.
 - Print in HTML, or a custom format.

## Usage

```go
package main

import (
	"fmt"

	"vimagination.zapto.org/pagination"
)

func main() {
	cfg := pagination.New()
	cfg.Ends = 2
	cfg.Surrounding = 2

	p := cfg.Get(5, 19)

	fmt.Println(p.HTML("/page?"))

	// Output:
	// <a href="/page?1">1</a> <a href="/page?2">2</a> <a href="/page?3">3</a> <a href="/page?4">4</a> <a href="/page?5">5</a> 6 <a href="/page?7">7</a> <a href="/page?8">8</a> ...<a href="/page?19">19</a> <a href="/page?20">20</a>
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/pagination
