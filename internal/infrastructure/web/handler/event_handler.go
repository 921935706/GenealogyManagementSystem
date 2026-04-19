package handler

import (
	"GenealogyManagementSystem/internal/application/event"
	eventDomain "GenealogyManagementSystem/internal/domain/event"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EventHandler 事件处理器
type EventHandler struct {
	appService *event.ApplicationService
}

// NewEventHandler 创建事件处理器实例
func NewEventHandler(appService *event.ApplicationService) *EventHandler {
	return &EventHandler{appService: appService}
}

// Create 创建事件
// @Summary 创建事件
// @Description 创建新的家族事件
// @Tags 事件管理
// @Accept json
// @Produce json
// @Param event body eventDomain.Event true "事件信息"
// @Success 201 {object} eventDomain.Event
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event/create [post]
func (h *EventHandler) Create(c *gin.Context) {
	var e eventDomain.Event
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.CreateEvent(&e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, e)
}

// Update 更新事件
// @Summary 更新事件
// @Description 更新家族事件信息
// @Tags 事件管理
// @Accept json
// @Produce json
// @Param event body eventDomain.Event true "事件信息"
// @Success 200 {object} eventDomain.Event
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event/update [put]
func (h *EventHandler) Update(c *gin.Context) {
	var e eventDomain.Event
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.UpdateEvent(&e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

// Delete 删除事件
// @Summary 删除事件
// @Description 删除家族事件
// @Tags 事件管理
// @Produce json
// @Param id query int true "事件ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event/delete [delete]
func (h *EventHandler) Delete(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.appService.DeleteEvent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID 根据ID获取事件
// @Summary 根据ID获取事件
// @Description 根据ID获取家族事件信息
// @Tags 事件管理
// @Produce json
// @Param id query int true "事件ID"
// @Success 200 {object} eventDomain.Event
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event/getById [get]
func (h *EventHandler) GetByID(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	e, err := h.appService.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

// GetByPersonID 根据人员ID获取事件
// @Summary 根据人员ID获取事件
// @Description 根据人员ID获取其所有事件
// @Tags 事件管理
// @Produce json
// @Param personId query int true "人员ID"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event/getByPersonId [get]
func (h *EventHandler) GetByPersonID(c *gin.Context) {
	personIDStr := c.Query("personId")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	personID, err := strconv.Atoi(personIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	events, total, err := h.appService.GetEventsByPersonID(personID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  events,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetByType 根据事件类型获取事件
// @Summary 根据事件类型获取事件
// @Description 根据事件类型获取所有事件
// @Tags 事件管理
// @Produce json
// @Param type query string true "事件类型"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event/getByType [get]
func (h *EventHandler) GetByType(c *gin.Context) {
	eventType := c.Query("type")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	if eventType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event type"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	events, total, err := h.appService.GetEventsByType(eventType, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  events,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetImportant 获取重要事件
// @Summary 获取重要事件
// @Description 获取所有重要的家族事件
// @Tags 事件管理
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /event/getImportant [get]
func (h *EventHandler) GetImportant(c *gin.Context) {
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

	events, total, err := h.appService.GetImportantEvents(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  events,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}
