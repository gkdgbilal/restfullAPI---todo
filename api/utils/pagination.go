package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"strings"
)

// echo -> golang framework
// Pageable defined for pagination
type Pageable struct {
	Page          int64
	Size          int64
	Sort          []string `validate:"required,min=2,max=2"`
	Q             []string
	TotalElements int64
}

func NewPagination() *Pageable {
	return &Pageable{Page: 0, Size: 20, Sort: []string{"createdAt", "desc"}, Q: []string{}}
}

func (p *Pageable) Filter() []bson.E {
	var elements []bson.E
	for _, filter := range p.Q {
		var e primitive.E
		filters := strings.Split(filter, ":")
		if len(filters) < 2 || len(filters[0]) == 0 || len(filters[1]) == 0 {
			continue
		}
		key := ""
		value := ""
		sep := "."
		length := len(filters)
		for i, f := range filters {
			if length == i+2 {
				sep = ""
			}

			if length == i+1 {
				value = f
			} else {
				key += strcase.ToSnake(f) + sep
			}
			e = primitive.E{Key: key, Value: value}
		}
		elements = append(elements, e)
	}
	return elements
}

func (p *Pageable) IsLast() bool {
	return p.Page == p.GetTotalPage()
}

func (p *Pageable) HasNext() bool {
	return p.Page < p.GetTotalPage()-1
}

func (p *Pageable) GetTotalPage() int64 {
	totalPage := float64(p.TotalElements / p.Size)
	mod := p.TotalElements % p.Size
	if mod == 0 && p.TotalElements != 0 {
		totalPage -= 1
	}
	return int64(math.Ceil(totalPage))
}

func (p *Pageable) GetSortKey() string {
	return strcase.ToSnake(p.Sort[0])
}

func (p *Pageable) GetSortValue() int {
	switch p.Sort[1] {
	case "asc":
		return 1
	case "desc":
		return -1
	default:
		return -1
	}
}

const (
	HeaderXTotalCount = "X-Total-Count"
	HeaderLink        = "Link"
)

func (p *Pageable) PaginationHeader(ctx *fiber.Ctx) (int64, string) {
	host := ctx.BaseURL()
	path := strings.Split(ctx.OriginalURL(), "?")
	url := host + path[0]
	link := ""

	totalPage := p.TotalElements / p.Size
	if p.TotalElements%p.Size > 0 {
		totalPage += 1
	}
	if p.Page < totalPage-1 {
		link += fmt.Sprintf("%s,", prepareLink(url, p.Page+1, p, "next"))
	}

	if p.Page > 0 {
		link += fmt.Sprintf("%s, ", prepareLink(url, p.Page-1, p, "prev"))
	}

	link += fmt.Sprintf("%s, ", prepareLink(url, totalPage, p, "last"))

	link += fmt.Sprintf("%s", prepareLink(url, 0, p, "first"))

	return p.TotalElements, link
}

func prepareLink(host string, page int64, p *Pageable, relType string) string {
	return fmt.Sprintf("<%s?page=%d&size=%d&sort=%s,%s>; rel=\"%s\"", host, page, p.Size, p.Sort[0], p.Sort[1], relType)
}

// ?page=0&size=10&sort=id,desc
