package v1

import (
	"github.com/gin-gonic/gin"
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

func (h *Handler) Register(e *gin.Engine) {
	auth := e.Group("auth")
	{
		auth.POST("/sign-up", h.singnUp)
		auth.POST("/sign-in", h.singnIn)
	}

	api := e.Group("api", h.userIdentity)
	{
		api.GET("/users", h.getUsers) // admin

		e.GET("/result-test", h.getResultTest) // student
		e.POST("/result-test", h.checkTest)    // student
		//
		//e.GET("/theory", h.theory)
		//e.GET("/test", h.test)
	}
}

func (h *Handler) checkTest() {
	id, _ := e.Get(userCtx)

	res, err := h.service.CheckTest(id.(int64))
}

func (h *Handler) getResultTest(e *gin.Context) {
	id, _ := e.Get(userCtx)
	res, err := h.service.GetResultTests(id.(int64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	e.JSON(http.StatusOK, res)
}

//func (h *Handler) test(c echo.Context) error {
//
//}
//
//func (h *Handler) theory(c echo.Context) error {
//	result, _ := h.service.GetTheory()
//	return c.JSON(http.StatusOK, result)
//}

func (h *Handler) getUsers(e *gin.Context) {
	role, _ := e.Get(userRole)
	if role.(string) != "admin" {
		e.AbortWithStatusJSON(http.StatusForbidden, "Access Denied")
		return
	}

	users, err := h.service.GetAllUsers()
	if err != nil {
		e.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	e.JSON(http.StatusOK, users)
}

func (h *Handler) singnUp(c *gin.Context) {
	var req entity.User
	if err := c.Bind(&req); err != nil {
		logrus.Error("invalid input body")
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid input body")
		return
	}
	user, err := h.service.CreateUser(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) singnIn(c *gin.Context) {
	var req Input

	if err := c.Bind(&req); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(req.Name, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}
