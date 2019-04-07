package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"services/server/media/service"
	"services/server/user/middleware"
	"services/server/core/render"
	"services/server/core/util"
)

func List(c *gin.Context) {
	payload := c.MustGet("JWT_PAYLOAD").(*middleware.AuthPayload)

	if payload == nil {
		util.Redirect(c, "/")
		return
	}

	// fetch requested media info for given page
	data, err := service.GetMediaDataByUserId(payload.ID, util.GetSelectPage(c))
	if err != nil {
		render.RenderPage(c, http.StatusInternalServerError, nil)
		return
	}

	props := service.GetListMediaResponseProps(c, data)

	render.RenderPage(c, http.StatusOK, props)
}
