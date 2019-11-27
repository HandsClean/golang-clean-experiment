package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/handsclean/webservice-clean/models"
	"github.com/handsclean/webservice-clean/orders"
)

type mssqlOrdersRepo struct {
	DB *sql.DB
}

// NewMssqlOrderRepository will create an implementation of author.Repository
func NewMssqlOrderRepository(db *sql.DB) orders.Repository {
	return &mssqlOrdersRepo{
		DB: db,
	}
}

// ID        int64  `json:"id"`
// 	number1c      string `json:"number1c"`
// 	CreatedAt string `json:"created_at"`
// 	UpdatedAt string `json:"updated_at"
func (m *mssqlOrdersRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Order, error) {

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.Order{}

	err = row.Scan(
		&a.ID,
		&a.number1c,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlOrdersRepo) GetByID(ctx context.Context, id int64) (*models.Order, error) {
	query := `SELECT id, name, created_at, updated_at FROM Order WHERE id=?`
	return m.getOne(ctx, query, id)
}
