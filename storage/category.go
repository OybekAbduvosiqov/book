package storage

import (
	"database/sql"
	"fmt"
	"github.com/OybekAbduvosiqov/book/models"

	"github.com/google/uuid"
)

func InsertCategory(db *sql.DB, category models.CreateCategory) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO categorys (
			id,
			name
		) VALUES ($1, $2)
	`

	_, err := db.Exec(query,
		id,
		category.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetByIdCategory(db *sql.DB, req models.CategoryPrimeryKey) (models.Category, error) {

	var (
		category models.Category
	)

	query := `
		SELECT
			id,
			name
		FROM categorys WHERE id = $1
	`

	err := db.QueryRow(query, req.Id).Scan(
		&category.Id,
		&category.Name,
	)

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func GetListCategory(db *sql.DB, req models.GetListCategoryRequest) (models.GetListCategoryResponse, error) {

	var (
		resp   models.GetListCategoryResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name
		FROM categorys
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := db.Query(query)
	if err != nil {
		return models.GetListCategoryResponse{}, err
	}

	for rows.Next() {
		var category models.Category

		err = rows.Scan(
			&resp.Count,
			&category.Id,
		)

		if err != nil {
			return models.GetListCategoryResponse{}, err
		}

		resp.Categorys = append(resp.Categorys, category)
	}

	return resp, nil
}

func UpdateCategory(db *sql.DB, category models.UpdateCategory) error {

	query := `
		UPDATE 
			categorys 
		SET 
			name = $2,
		WHERE id = $1
	`

	_, err := db.Exec(query,
		category.Id,
		category.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM categorys WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
