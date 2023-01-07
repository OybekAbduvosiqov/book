package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/OybekAbduvosiqov/book/api/handler"
)

func NewApi(r *gin.Engine, db *sql.DB) {

	handlerV1 := handler.NewHandler(db)

	r.POST("/book", handlerV1.CreateBook)
	r.GET("/book/:id", handlerV1.GetByIDBook)
	r.GET("/book", handlerV1.GetListBook)
	r.DELETE("/book/:id", handlerV1.DeleteBook)
	r.PUT("/book", handlerV1.UpdateBook)

	r.POST("/category", handlerV1.CreateCategory)
	r.GET("/category/:id", handlerV1.GetByIDCategory)
	r.GET("/category", handlerV1.GetListCategory)

}
