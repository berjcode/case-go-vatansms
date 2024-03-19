package handlers

import "github.com/labstack/echo/v4"

type StaticHandler struct{}

func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) IndexHTML(c echo.Context) error {
	return c.File("templates/index.html")
}
