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
			name,
			updated_at
		) VALUES ($1, $2, now())
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

func GetByIdCategory(db *sql.DB, req models.CategoryPrimeryKey) (models.CategoryB, error) {

	var (
		Category models.CategoryB
	)
	query := `
		select
			categorys.id,
			categorys.name
		from book_category
		join categorys on categorys.id = book_category.categorys_id
		where book_category.categorys_id = $1
	`

	query1 := `
		select
			books.id,
			books.name,
			books.price,
			books.description
		from book_category
		join books on books.id = book_category.books_id
		where book_category.categorys_id = $1
		group by books.name, books.id;
	`
	rows, err := db.Query(query, req.Id)
	if err != nil {
		return models.CategoryB{}, err
	}

	for rows.Next() {
		err = rows.Scan(
			&Category.Id,
			&Category.Name,
		)
		rows, err = db.Query(query1, req.Id)
		for rows.Next() {
			var book models.BookInfo
			err = rows.Scan(
				&book.Id,
				&book.Name,
				&book.Price,
				&book.Description,
			)
			Category.BookInfos = append(Category.BookInfos, book)
		}
	}
	if err != nil {
		return models.CategoryB{}, err
	}

	return Category, nil
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
			name,
			created_at,
			updated_at
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
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return models.GetListCategoryResponse{}, err
		}

		resp.Categorys = append(resp.Categorys, category)
	}

	return resp, nil
}

func UpdateCategory(db *sql.DB, category models.UpdateCategory) (int64, error) {

	query := `
		UPDATE 
			categorys 
		SET 
			name = $2,
			updated_at = now()
		WHERE id = $1
	`

	result, err := db.Exec(query,
		category.Id,
		category.Name,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteCategory(db *sql.DB, req models.CategoryPrimeryKey) error {
	_, err := db.Exec("DELETE FROM categorys WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}
