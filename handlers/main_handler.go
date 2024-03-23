package handlers

import (
	"github.com/labstack/echo/v4"
)

type StaticHandler struct{}

const (
	IndexHTMLPath        = "templates/index.html"
	RegisterHTMLPath     = "templates/register.html"
	PlanPageHTMLPath     = "templates/plan.html"
	UserDetailPagePath   = "templates/userdetail.html"
	LessonPageHTMLPath   = "templates/lesson.html"
	LessonDetailPageHtml = "templates/lessondetail.html"
)

func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) IndexHTML(c echo.Context) error {
	return c.File(IndexHTMLPath)
}

func (h *StaticHandler) RegisterHTML(c echo.Context) error {
	return c.File(RegisterHTMLPath)
}

func (h *StaticHandler) PlanPageHTML(c echo.Context) error {
	return c.File(PlanPageHTMLPath)
}

func (h *StaticHandler) LoginPageHtml(c echo.Context) error {
	return c.File(IndexHTMLPath)
}

func (h *StaticHandler) UserDetailPage(c echo.Context) error {
	return c.File(UserDetailPagePath)
}

func (h *StaticHandler) LessonPageHtml(c echo.Context) error {
	return c.File(LessonPageHTMLPath)
}

func (h *StaticHandler) LessonDetailPageHtml(c echo.Context) error {
	return c.File(LessonDetailPageHtml)
}
