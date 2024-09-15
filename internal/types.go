package internal

import (
	"strconv"
	"strings"
)

type URLParams struct {
	Title          string
	Director       string
	Producers      []string
	Prod_companies []string
	Writers        []string
	Cast_members   []string
	Genres         []string
	Age_rating     string
	Status         string
	Country        string
	Sort           string
	Print          string
	Page           int
	Page_size      int
}

func (p *URLParams) SortColumn() string {
	return strings.TrimPrefix(p.Sort, "-")
}

func (p *URLParams) SortDirection() string {
	if strings.HasPrefix(p.Sort, "-") {
		return "DESC"
	}

	return "ASC"
}

func (p *URLParams) Limit() string {
	if p.Page_size != 0 {
		return strconv.Itoa(p.Page_size)
	} else {
		return "ALL"
	}
}

func (p *URLParams) Offset() int {
	return (p.Page - 1) * p.Page_size
}
