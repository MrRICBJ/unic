package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"university/internal/entity"
	"university/internal/service"
)

type Handler struct {
	service *service.Service
}

type Input struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(e *echo.Echo) {
	auth := e.Group("auth")
	{
		auth.POST("/sign-up", h.singnUp)
		auth.POST("/sign-in", h.singnIn)
	}

	e.GET("/users", h.getUsers)
	e.GET("/theory", h.theory)
	e.GET("/test", h.test)
}

func (h *Handler) test(c echo.Context) error {

}

func (h *Handler) theory(c echo.Context) error {
	result, _ := h.service.GetTheory()
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) singnUp(c echo.Context) error {
	var req entity.User
	if err := c.Bind(&req); err != nil {
		logrus.Error("invalid input body")
		return c.JSON(http.StatusBadRequest, "invalid input body")
	}
	user, err := h.service.CreateUser(&req)
	//id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) singnIn(c echo.Context) error {
	var req Input

	if err := c.Bind(&req); err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.GenerateToken(req.Name, req.Password)
	//id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Handler) getUsers(e echo.Context) error {
	return nil
}
