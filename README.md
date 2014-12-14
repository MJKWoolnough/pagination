# pagination
--
    import "github.com/MJKWoolnough/pagination"

Package pagination implements a pagination solution for multiple front-ends

## Usage

#### type Config

```go
type Config struct {
	Ends, Surrounding uint
}
```

Config is the configuation for a Pagination

#### func  New

```go
func New() Config
```
New returns a default configuration for Pagination

#### func (Config) Get

```go
func (c Config) Get(currPage, lastPage uint) Pagination
```
Get returns the Section information for Pagination

#### type Pagination

```go
type Pagination struct {
}
```

Pagination contains the information necessary to print a proper pagination

#### func (Pagination) HTML

```go
func (p Pagination) HTML(urlBase string) string
```
HTML calls Print with a HTML based pageFn and a simple elipses.

The urlBase will have the page number appended to it, so it needs to be
formatted with this in mind.

#### func (Pagination) Print

```go
func (p Pagination) Print(pageFn func(uint) string, sep string) string
```
Print converts the Pagination sections into a string.

The pageFn func takes a page number and returns whatever text is needed for that
page.

The sep string is what is to appear between the sections.

#### func (Pagination) String

```go
func (p Pagination) String() string
```
String stringifies the Sections with a simple pageFn
