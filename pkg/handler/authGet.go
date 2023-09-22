package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTokens(c *gin.Context) {
	GUID, err := strconv.Atoi(c.Params.ByName("GUID"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	ShowID, err := h.services.Authorization.GetTokens(GUID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Id form http": ShowID,
	})
}
