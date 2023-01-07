package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/OybekAbduvosiqov/book/models"
	"github.com/OybekAbduvosiqov/book/storage"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategory(c *gin.Context) {

	var category models.CreateCategory

	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.InsertCategory(h.db, category)
	if err != nil {
		log.Println("error whiling create category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := storage.GetByIdCategory(h.db, models.CategoryPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) GetByIDCategory(c *gin.Context) {

	id := c.Param("id")

	res, err := storage.GetByIdCategory(h.db, models.CategoryPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) GetListCategory(c *gin.Context) {
	var (
		err       error
		offset    int
		limit     int
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")
	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Println("error whiling offset:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("error whiling limit:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res, err := storage.GetListCategory(h.db, models.GetListCategoryRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}
