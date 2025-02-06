package db

import (
	"context"

	"github.com/uptrace/bun"
)

var db = InitDBConnect()

func Create(model interface{}) (q *bun.InsertQuery, ctx context.Context) {
	ctx = context.Background()

	q = db.NewInsert().Model(model)
	return
}

func Read(model interface{}) (q *bun.SelectQuery, ctx context.Context) {
	ctx = context.Background()

	q = db.NewSelect().Model(model)
	return
}

func Update(model interface{}) (q *bun.UpdateQuery, ctx context.Context) {
	ctx = context.Background()

	q = db.NewUpdate().Model(model)
	return
}

func Delete(model interface{}) (q *bun.DeleteQuery, ctx context.Context) {
	ctx = context.Background()

	q = db.NewDelete().Model(model)
	return
}
