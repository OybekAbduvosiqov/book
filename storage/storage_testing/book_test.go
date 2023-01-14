package storage_testing

import (
	"context"
	"testing"

	"github.com/OybekAbduvosiqov/book/models"
	"github.com/google/go-cmp/cmp"
)

func TestBookInsert(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateBook
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CreateBook{
				Name:        "Time",
				Price:       22000,
				Description: "OK",
			},
			WantErr: false,
		},
		{
			Name: "case 2",
			Input: &models.CreateBook{
				Price:       22000,
				Description: "OK",
			},
			WantErr: false,
		},
		{
			Name: "case 3",
			Input: &models.CreateBook{
				Price:       22000,
				Description: "OK",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			got, err := BookRepo.Insert(context.Background(), tc.Input)
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}
			if got == "" {
				t.Errorf("%s: got: %v", tc.Name, got)
				return
			}
		})
	}
}

func TestBookGetById(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.BookPrimeryKey
		Output  *models.Book
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.BookPrimeryKey{
				Id: "9ef1a045-1e2a-41f5-8c5d-43f9a23d7cfd",
			},
			Output: &models.Book{
				Name:        "Time",
				Price:       23000,
				Description: "OK",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			got, err := BookRepo.GetByID(context.Background(), tc.Input)
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Book) bool {
				return x.Name == y.Name &&
					x.Price == y.Price &&
					x.Description == y.Description
			})

			if diff := cmp.Diff(tc.Output, got, comparer); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func TestUpdate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.BookPrimeryKey
		Output  *models.Book
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.BookPrimeryKey{
				Id: "9ef1a045-1e2a-41f5-8c5d-43f9a23d7cfd",
			},
			Output: &models.Book{
				Name:        "Time123",
				Price:       23000,
				Description: "OK",
				UpdatedAt:   "2023-10-10",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			got, err := BookRepo.Update(context.Background())
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Book) bool {
				return x.Name == y.Name &&
					x.Price == y.Price &&
					x.Description == y.Description &&
					x.UpdatedAt == y.UpdatedAt
			})

			if diff := cmp.Diff(tc.Output, got, comparer); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func TestDelete(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.BookPrimeryKey
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.BookPrimeryKey{
				Id: "9ef1a045-1e2a-41f5-8c5d-43f9a23d7cfd",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			got, err := BookRepo.Delete(context.Background())
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}
		})
	}
}
