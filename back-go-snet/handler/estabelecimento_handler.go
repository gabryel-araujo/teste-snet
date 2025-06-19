package handler

import (
	"database/sql"
	"net/http"

	// "strconv"

	"github.com/gabryel-araujo/back-go-snet/models"
	"github.com/gabryel-araujo/back-go-snet/service"
	"github.com/labstack/echo/v4"
)

type EstabHandler struct{
	Service *service.EstabService
}

//POST
func (h *EstabHandler) CreateEstab(c echo.Context) error {
	e := new(models.Estabelecimentos)

	if err := c.Bind(e); err != nil {
		return c.JSON(http.StatusBadRequest,err)
	}

	err := h.Service.CreateEstab(e)

	if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{
        "error": "Erro ao criar estabelecimento",
        "details": err.Error(), 
    })
	}
	return c.JSON(http.StatusCreated,e)
}

//GET
func (h *EstabHandler) GetAllEstabs(c echo.Context) error {
	estabs, err := h.Service.GetAllEstabs()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao buscar estabelecimentos"})
	}
	
	return c.JSON(http.StatusOK, estabs)
}

//GET/id
func (h *EstabHandler) GetOneEstab(c echo.Context) error {
	id := c.Param("id")

	estabs,err := h.Service.GetOneEstab(id)

	if err == sql.ErrNoRows{
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Estabelecimento não encontrado"})
	}

	return c.JSON(http.StatusOK,estabs)
}

//PUT/id
func (h *EstabHandler) UpdateEstab(c echo.Context) error {
	id := c.Param("id")
	e := new(models.Estabelecimentos)

	if err := c.Bind(e); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	if err := h.Service.UpdateEstab(id, e); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Estabelecimento atualizado com sucesso"})
}

//DELETE/id
func (h *EstabHandler) DeleteEstab(c echo.Context) error {
	id := c.Param("id")

	if err := h.Service.DeleteEstab(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Estabelecimento removido com sucesso"})
}
