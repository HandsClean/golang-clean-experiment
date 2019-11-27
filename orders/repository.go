package orders

import (
	"context"

	"github.com/handsclean/webservice-clean/models"
)

// Repository represent the author's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.Order, error)
}
