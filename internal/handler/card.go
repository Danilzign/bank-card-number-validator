package handler

import (
	"net/http"
	"strconv"
	"validate"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createCard(c *gin.Context) {

	number := c.PostForm("number")
	digit, _ := strconv.Atoi(number)

	var input validate.Card
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Card.Create(digit, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"1id": id,
	})

}

type getAllCardsResponse struct {
	Data []validate.Card `json:"data"`
}

func (h *Handler) getAllCards(c *gin.Context) {
	limit := c.Request.URL.Query().Get("limit")
	page := c.Request.URL.Query().Get("page")

	products, err := h.services.Card.GetAll(limit, page)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCardsResponse{
		Data: products,
	})

}

func (h *Handler) getCardById(c *gin.Context) {
	id := c.Param("id")

	product, err := h.services.Card.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) updateCard(c *gin.Context) {
	id := c.Param("id")

	var input validate.UpdateCardInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Card.UpdateCard(id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteCard(c *gin.Context) {

	id := c.Param("id")

	err := h.services.Card.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
