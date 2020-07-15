package v1

import (
	"database/sql"
	"net/http"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/responses"

	"github.com/labstack/echo/v4"
)

func (h Handlers) GetGuildList(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	guildList, err := h.DB.GetGuildList(ctx.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, responses.MetadataNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	return ctx.JSON(http.StatusOK, GetGuildListResponse{
		Guilds: guildList,
	})
}