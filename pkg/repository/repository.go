package repository

import (
	"time"

	"github.com/jinyanomura/ezres-web/pkg/models"
)

type DatabaseRepo interface {
	GetAllTables() ([]models.Table, error)
	GetRestrictionsByDay(int, time.Time) ([]models.Restriction, error)
}