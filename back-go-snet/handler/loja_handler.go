package handler

import (
	"net/http"

	"github.com/gabryel-araujo/back-go-snet/models"
	"github.com/gabryel-araujo/back-go-snet/service"
	"github.com/labstack/echo/v4"
)

type StoreHandler struct{
	Service *service.StoreService
}

//POST
func (h *StoreHandler) CreateStore(c echo.Context) error {
	l := new(models.Lojas)

	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest,err)
	}

	err := h.Service.CreateStore(l)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Erro ao criar loja",
			"details": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated,l)
}

//GET
func (h *StoreHandler) GetAllStores(c echo.Context) error {
	stores,err := h.Service.Repo.GetAllStores()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao buscar lojas"})
	}

	return c.JSON(http.StatusOK, stores)
}

//GET/id
func (h *StoreHandler) GetOneStore(c echo.Context) error {
	id := c.Param("id")

	store, err := h.Service.GetOneStore(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK,store)
}

//PUT/id
func (h *StoreHandler) UpdateStore(c echo.Context) error {
	id := c.Param("id")
	l := new(models.Lojas)

	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	if err := h.Service.UpdateStore(id, l); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Loja atualizada com sucesso"})
}


//DELETE/id
func (h *StoreHandler) DeleteStore(c echo.Context) error {
	id := c.Param("id")

	if err := h.Service.DeleteStore(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Loja removida com sucesso"})
}
