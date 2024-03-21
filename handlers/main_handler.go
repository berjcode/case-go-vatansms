package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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
func (h *StaticHandler) LoginPageHtml(c echo.Context) error {
	return c.File("templates/index.html")
}
func (h *StaticHandler) UserDetailPage(c echo.Context) error {
	return c.File("templates/userdetail.html")
}

func (h *StaticHandler) UserDetailData(c echo.Context) error {
	user, err := GetUserByUsername(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *StaticHandler) LessonPageHtml(c echo.Context) error {
	return c.File("templates/lesson.html")
}
