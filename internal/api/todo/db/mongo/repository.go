package mongo

import (
	"todo/internal/api/todo/db"
)

type repository struct{}

func GetRepository() db.IRepository {
	return &repository{}
}
