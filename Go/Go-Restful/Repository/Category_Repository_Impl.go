package Repository

import (
	"context"
	"database/sql"
	"mandala/restful-api/Helper"
	"mandala/restful-api/Model/Domain"
)

type CategoryRepositoryImpl struct {
}

func (repository CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category Domain.Category) Domain.Category {
	sql := "insert into category(name) values (?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	Helper.PanicIfError(err)

	id, err := result.LastInsertId()
	Helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category Domain.Category) CategoryRepository {
	//TODO implement me
	panic("implement me")
}

func (repository CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category CategoryRepository) {
	//TODO implement me
	panic("implement me")
}

func (repository CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) CategoryRepository {
	//TODO implement me
	panic("implement me")
}

func (repository CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []CategoryRepository {
	//TODO implement me
	panic("implement me")
}
