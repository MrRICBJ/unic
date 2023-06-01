package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"university/internal/entity"
	"university/internal/handlers"
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
		api.GET("/users", h.getUsers) // admin - ok

		api.GET("/result-test", h.getResultTest) // student -ok
		api.POST("/result-test", h.checkTest)    // student - ok

		api.POST("/add-student", h.addStudent)   //teacher - ok
		api.GET("/my-students", h.getMyStudents) //teacher -
	}

	e.GET("/theory", h.theory)
	e.GET("/testQuestions", h.testQuestions)
	e.GET("/testAnswers", h.testAnswers)
}

func (h *Handler) theory(e *gin.Context) {
	theory, _ := h.service.GetTheory()
	e.JSON(http.StatusOK, theory)
}

func (h *Handler) testQuestions(e *gin.Context) {
	res, _ := h.service.GetTestQuestions()
	e.JSON(http.StatusOK, res)
}

func (h *Handler) testAnswers(e *gin.Context) {
	res, _ := h.service.GetTestAnswers()
	e.JSON(http.StatusOK, res)
}

func (h *Handler) getMyStudents(e *gin.Context) {
	id, _ := e.Get(userCtx)

	res, err := h.service.GetMyStudent(id.(int64))
	if err != nil {
		e.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	e.JSON(http.StatusOK, res)
}

func (h *Handler) addStudent(e *gin.Context) {
	id, _ := e.Get(userCtx)
	idStudent, err := strconv.Atoi(e.Query("idStudent"))
	if err != nil {
		e.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.AddStudent(idStudent, id.(int64)); err != nil {
		e.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	e.JSON(http.StatusOK, idStudent)
}

func (h *Handler) checkTest(e *gin.Context) {
	id, _ := e.Get(userCtx)
	var req handlers.RequestBodyTest
	err := e.BindJSON(&req)
	if err != nil {
		e.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.service.CheckTest(id.(int64), req.Answer)
	if err != nil {
		e.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	e.JSON(http.StatusOK, fmt.Sprintf("%d/13", res))
}

func (h *Handler) getResultTest(e *gin.Context) {
	id, _ := e.Get(userCtx)
	res, err := h.service.GetResultTests(id.(int64))
	if err != nil {
		e.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	e.JSON(http.StatusOK, res)
}

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
