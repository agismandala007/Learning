package Repository

import (
	"context"
	"database/sql"
	"mandala/restful-api/Model/Domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category Domain.Category) Domain.Category
	Update(ctx context.Context, tx *sql.Tx, category Domain.Category) CategoryRepository
	Delete(ctx context.Context, tx *sql.Tx, category CategoryRepository)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) CategoryRepository
	FindAll(ctx context.Context, tx *sql.Tx) []CategoryRepository
}
