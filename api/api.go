package api

import (
	// "database/sql"

	"github.com/OybekAbduvosiqov/book/storage"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	_ "github.com/OybekAbduvosiqov/book/api/docs"
	"github.com/OybekAbduvosiqov/book/api/handler"
)

func NewApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandler(storage)

	r.POST("/book", handlerV1.CreateBook)
	r.GET("/book/:id", handlerV1.GetByIDBook)
	r.GET("/book", handlerV1.GetListBook)
	r.PUT("/book/:id", handlerV1.UpdateBook)
	r.DELETE("/book/:id", handlerV1.DeleteBook)

	r.POST("/category", handlerV1.CreateCategory)
	r.GET("/category/:id", handlerV1.GetByIdCategory)
	r.GET("/category", handlerV1.GetListCategory)
	r.DELETE("/category/:id", handlerV1.DeleteCategory)
	r.PUT("/category/:id", handlerV1.UpdateCategory)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
