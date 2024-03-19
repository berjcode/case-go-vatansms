package handlers

import "github.com/labstack/echo/v4"

type StaticHandler struct{}

func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) IndexHTML(c echo.Context) error {
	return c.File("templates/index.html")
}

func (h *StaticHandler) RegisterHTML(c echo.Context) error {
	return c.File("templates/register.html")
}

func (h *StaticHandler) PlanPageHTML(c echo.Context) error {
	return c.File("templates/plan.html")
}
