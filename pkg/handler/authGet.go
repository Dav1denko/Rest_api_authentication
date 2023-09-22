package handler

import (
	"fmt"
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
	Response, err := h.services.Authorization.GetTokens(GUID)
	if err != nil {
		return
	}

	CookieGUID, err := c.Cookie("GUID")
	if err != nil {
		fmt.Sprintln("error get cookie")
	}

	GUID_id := strconv.Itoa(Response.GUID)

	if CookieGUID == GUID_id {
		c.JSON(http.StatusOK, map[string]interface{}{
			"Error": "GUID has JWT",
			"GUID":  GUID,
		})
	} else {
		h.services.SaveRefreshToken(Response.GUID, Response.RefreshToken)
		c.SetCookie("refresh-token", Response.RefreshToken, 3600, "", "", true, true)
		c.SetCookie("GUID", GUID_id, 3600, "", "", true, true)
		c.JSON(http.StatusOK,
			Response,
		)
	}

}
