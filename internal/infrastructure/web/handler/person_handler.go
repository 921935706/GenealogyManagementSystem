package handler

import (
	"GenealogyManagementSystem/internal/application/person"
	personDomain "GenealogyManagementSystem/internal/domain/person"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PersonHandler 人员处理器
type PersonHandler struct {
	appService *person.ApplicationService
}

// NewPersonHandler 创建人员处理器实例
func NewPersonHandler(appService *person.ApplicationService) *PersonHandler {
	return &PersonHandler{appService: appService}
}

// Create 创建人员
// @Summary 创建人员
// @Description 创建新的家族成员
// @Tags 人员管理
// @Accept json
// @Produce json
// @Param person body personDomain.Person true "人员信息"
// @Success 201 {object} personDomain.Person
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /person/create [post]
func (h *PersonHandler) Create(c *gin.Context) {
	var p personDomain.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.CreatePerson(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

// Update 更新人员信息
// @Summary 更新人员信息
// @Description 更新家族成员的信息
// @Tags 人员管理
// @Accept json
// @Produce json
// @Param person body personDomain.Person true "人员信息"
// @Success 200 {object} personDomain.Person
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /person/update [put]
func (h *PersonHandler) Update(c *gin.Context) {
	var p personDomain.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.UpdatePerson(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

// Delete 删除人员
// @Summary 删除人员
// @Description 删除家族成员
// @Tags 人员管理
// @Produce json
// @Param id query int true "人员ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /person/delete [delete]
func (h *PersonHandler) Delete(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.appService.DeletePerson(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID 根据ID获取人员
// @Summary 根据ID获取人员
// @Description 根据ID获取家族成员信息
// @Tags 人员管理
// @Produce json
// @Param id query int true "人员ID"
// @Success 200 {object} personDomain.Person
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /person/getById [get]
func (h *PersonHandler) GetByID(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	p, err := h.appService.GetPersonByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

// GetByUUID 根据UUID获取人员
// @Summary 根据UUID获取人员
// @Description 根据UUID获取家族成员信息
// @Tags 人员管理
// @Produce json
// @Param uuid query string true "人员UUID"
// @Success 200 {object} personDomain.Person
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /person/getByUuid [get]
func (h *PersonHandler) GetByUUID(c *gin.Context) {
	uuid := c.Query("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	p, err := h.appService.GetPersonByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

// List 列出人员
// @Summary 列出人员
// @Description 分页列出家族成员
// @Tags 人员管理
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /person/list [get]
func (h *PersonHandler) List(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	persons, total, err := h.appService.ListPersons(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  persons,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// Search 搜索人员
// @Summary 搜索人员
// @Description 根据姓名搜索家族成员
// @Tags 人员管理
// @Produce json
// @Param name query string true "搜索姓名"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /person/search [get]
func (h *PersonHandler) Search(c *gin.Context) {
	name := c.Query("name")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	persons, total, err := h.appService.SearchPersons(name, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  persons,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}
