package models

import (
	"fmt"
	"strings"
	"time"

	base "app_store_api/db"
	"app_store_api/internal"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type (
	Products struct {
		Total    int        `json:"total"`
		Results  *[]Product `json:"results"`
		MaxPrice float64    `json:"max_price"`
		MinPrice float64    `json:"min_price"`
	}

	Product struct {
		bun.BaseModel `bun:"table:products,alias:p"`

		ID          int64     `bun:"id,pk,autoincrement" json:"id"`
		Name        string    `bun:"name" json:"name"`
		Description string    `bun:"description" json:"description"`
		Price       float64   `bun:"price" json:"price"`
		Category    string    `bun:"category" json:"category"`
		Status      string    `bun:"status,default:O" json:"status"`
		Flag        string    `bun:"flag,default:dev" json:"flag"`
		Active      bool      `bun:"active,default:true" json:"active"`
		UUID        string    `bun:"uuid,default:uuid_generate_v4()" json:"uuid"`
		CreatedAt   time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
		UpdatedAt   time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
		DeletedAt   time.Time `bun:"deleted_at,soft_delete,nullzero" json:"deleted_at"`
	}

	Price struct {
		MaxPrice float64 `bun:"max_price" json:"max_price"`
		MinPrice float64 `bun:"min_price" json:"min_price"`
	}
)

func (m *Product) ReadAllProducts(limit, offset int, keyword, qExt string, sortBy []string) (data *[]Product, total int, err error) {
	data = new([]Product)

	cols := []string{
		"name",
		"description",
	}

	var columns []string
	for _, col := range cols {
		coalesce := "coalesce(" + col + ",'')"
		columns = append(columns, coalesce)
	}

	q, ctx := base.Read(data)

	if keyword != "" {
		q = q.Where(fmt.Sprintf("%s ~* ?", strings.Join(columns, " || ")), keyword)
	}

	q = m.applyFilter(q, qExt)

	q = q.Limit(limit)
	q = q.Offset(offset)

	if len(sortBy) > 0 {
		for _, sort := range sortBy {
			q = q.Order(sort)
		}
	} else {
		q = q.Order("p.created_at DESC")
	}

	total, err = q.ScanAndCount(ctx)

	return
}

func (m *Product) ReadOneProduct(uuid string) (data *Product, err error) {
	data = new(Product)

	q, ctx := base.Read(data)

	q = q.Where("uuid = ?", uuid)

	err = q.Scan(ctx)

	return
}

func (m *Product) CreateProduct(data *Product) (*Product, error) {
	q, ctx := base.Create(data)

	_, err := q.Exec(ctx)

	data, err = m.cleanErr(data, err)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *Product) cleanErr(data *Product, err error) (*Product, error) {
	if pgErr, ok := err.(pgdriver.Error); ok {
		if pgErr.IntegrityViolation() {
			switch pgErr.Field('n') {
			case "product_name_unique_idx":
				return nil, internal.ErrNameUnique
			}
		}
	}

	return data, nil
}

func (m *Product) GetMinMaxPrice() (*Price, error) {
	price := new(Price)

	q, ctx := base.Read(new(Product))
	q = q.ColumnExpr("MAX(price) as max_price, MIN(price) as min_price")
	err := q.Scan(ctx, &price.MaxPrice, &price.MinPrice)

	if err != nil {
		return nil, err
	}

	return price, nil
}

func (m *Product) applyFilter(q *bun.SelectQuery, qExt string) *bun.SelectQuery {
	splitQExt := splitQExt(qExt)

	if len(splitQExt) > 0 {
		key := splitQExt[0]
		value := splitQExt[1]

		if key == "price" {
			splitValue := strings.Split(value, "-")
			q = q.Where("price >= ?", splitValue[0]).Where("price <= ?", splitValue[1])
		} else if key == "category" {
			splitValue := strings.Split(value, ",")
			q = q.Where("category IN (?)", bun.In(splitValue))
		} else {
			q = q.Where(fmt.Sprintf("%s = %s", key, value))
		}
	}

	return q
}
