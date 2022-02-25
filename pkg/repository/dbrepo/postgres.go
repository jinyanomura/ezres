package dbrepo

import (
	"context"
	"time"

	"github.com/jinyanomura/ezres-web/pkg/models"
)

func (m *postgresDBRepo) GetAllTables() ([]models.Table, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var tables []models.Table

	rows, err := m.DB.QueryContext(ctx, "select id, capacity from tables")
	if err != nil {
		return tables, err
	}
	defer rows.Close()

	for rows.Next() {
		var table models.Table
		err = rows.Scan(&table.ID, &table.Capacity)
		if err != nil {
			return tables, err
		}
		tables = append(tables, table)
	}

	if err = rows.Err(); err != nil {
		return tables, err
	}

	return tables, nil
}

func (m *postgresDBRepo) GetRestrictionsByDay(id int, date time.Time) ([]models.Restriction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var restrictions []models.Restriction

	query := `
		select id, start_time, end_time, table_id, restriction_id, reservation_id, created_at, updated_at
		from table_restrictions
		where table_id = $1 and start_time >= $2 and end_time < $3
	`

	rows, err := m.DB.QueryContext(ctx, query, id, date, date.AddDate(0, 0, 1))
	if err != nil {
		return restrictions, err
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Restriction
		err = rows.Scan(
			&r.ID,
			&r.StartTime,
			&r.EndTime,
			&r.TableID,
			&r.RestrictionID,
			&r.ReservationID,
			&r.CreatedAt,
			&r.UpdatedAt,
		)
		if err != nil {
			return restrictions, err
		}
		restrictions = append(restrictions, r)
	}

	if err = rows.Err(); err != nil {
		return restrictions, err
	}

	return restrictions, nil
}