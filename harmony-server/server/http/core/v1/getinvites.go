package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

type GetInvitesData struct {
	Guild int64 `validate:"required"`
}

func (h Handlers) GetInvites(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data GetInvitesData
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] == nil {
		return echo.NewHTTPError(http.StatusNotFound, "guild not found")
	}
	if isOwner, err := h.Deps.DB.IsOwner(data.Guild, ctx.UserID); err != nil || !isOwner {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permission to list invites")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many invite listing requests, please try again later")
	}
	invites, err := h.Deps.DB.GetInvites(data.Guild)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to get invites")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"invites": invites,
	})
}
