package models

import (
	base "app_store_api/db"
	"app_store_api/internal"
	"fmt"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type (
	Categories struct {
		Total   int         `json:"total"`
		Results *[]Category `json:"results"`
	}

	Category struct {
		bun.BaseModel `bun:"table:categories,alias:c"`

		ID          int64     `bun:"id,pk,autoincrement" json:"id"`
		Name        string    `bun:"name" json:"name"`
		Description string    `bun:"description" json:"description"`
		Status      string    `bun:"status,default:O" json:"status"`
		Flag        string    `bun:"flag,default:dev" json:"flag"`
		Active      bool      `bun:"active,default:true" json:"active"`
		UUID        string    `bun:"uuid,default:uuid_generate_v4()" json:"uuid"`
		CreatedAt   time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
		UpdatedAt   time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
		DeletedAt   time.Time `bun:"deleted_at,soft_delete,nullzero" json:"deleted_at"`
	}
)

func (m *Category) ReadAllCategories(limit, offset int, keyword, qExt string, sortBy []string) (data *[]Category, total int, err error) {
	data = new([]Category)

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
		q = q.Order("c.created_at DESC")
	}

	total, err = q.ScanAndCount(ctx)

	return
}

func (m *Category) ReadOneCategory(uuid string) (data *Category, err error) {
	data = new(Category)

	q, ctx := base.Read(data)

	q = q.Where("uuid = ?", uuid)

	err = q.Scan(ctx)

	return
}

func (m *Category) CreateCategory(data *Category) (*Category, error) {
	q, ctx := base.Create(data)

	_, err := q.Exec(ctx)

	data, err = m.cleanErr(data, err)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *Category) cleanErr(data *Category, err error) (*Category, error) {
	if pgErr, ok := err.(pgdriver.Error); ok {
		if pgErr.IntegrityViolation() {
			switch pgErr.Field('n') {
			case "category_name_unique_idx":
				return nil, internal.ErrNameUnique
			}
		}
	}

	return data, nil
}

func (m *Category) applyFilter(q *bun.SelectQuery, qExt string) *bun.SelectQuery {
	splitQExt := splitQExt(qExt)

	if len(splitQExt) > 0 {
		key := splitQExt[0]
		value := splitQExt[1]

		if key == "price" {
			splitValue := strings.Split(value, "-")
			q = q.Where("price >= ?", splitValue[0]).Where("price <= ?", splitValue[1])
		} else {
			q = q.Where(fmt.Sprintf("%s = %s", key, value))
		}
	}

	return q
}
