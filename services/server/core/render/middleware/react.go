package middleware

import (
	"sync"
	"github.com/gin-gonic/gin"
	"services/server/core/render"
)

var reactOnce sync.Once

func ReactRenderMiddleware(filePath string, debug bool, router *gin.Engine) (gin.HandlerFunc) {
	var react *render.React
	reactOnce.Do(func() {
		react = render.NewReact(filePath, debug, router)
	})
	return func(c *gin.Context) {
		c.Next() // handle the request

		// only render the response if react metadata was set (so we know that we want to render through react
		if _, exists := c.Get("react-meta"); exists {
			react.Handle(c)
		}
	}
}
