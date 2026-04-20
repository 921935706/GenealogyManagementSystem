package handler

import (
	"GenealogyManagementSystem/internal/application/media"
	mediaDomain "GenealogyManagementSystem/internal/domain/media"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MediaHandler 多媒体处理器
type MediaHandler struct {
	appService *media.ApplicationService
}

// NewMediaHandler 创建多媒体处理器实例
func NewMediaHandler(appService *media.ApplicationService) *MediaHandler {
	return &MediaHandler{appService: appService}
}

// Create 创建多媒体
// @Summary 创建多媒体
// @Description 创建新的家族多媒体文件
// @Tags 多媒体管理
// @Accept json
// @Produce json
// @Param media body mediaDomain.Media true "多媒体信息"
// @Success 201 {object} mediaDomain.Media
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /media/create [post]
func (h *MediaHandler) Create(c *gin.Context) {
	var m mediaDomain.Media
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.CreateMedia(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, m)
}

// Update 更新多媒体
// @Summary 更新多媒体
// @Description 更新家族多媒体文件信息
// @Tags 多媒体管理
// @Accept json
// @Produce json
// @Param media body mediaDomain.Media true "多媒体信息"
// @Success 200 {object} mediaDomain.Media
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /media/update [put]
func (h *MediaHandler) Update(c *gin.Context) {
	var m mediaDomain.Media
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.UpdateMedia(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, m)
}

// Delete 删除多媒体
// @Summary 删除多媒体
// @Description 删除家族多媒体文件
// @Tags 多媒体管理
// @Produce json
// @Param id query int true "多媒体ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /media/delete [delete]
func (h *MediaHandler) Delete(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.appService.DeleteMedia(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID 根据ID获取多媒体
// @Summary 根据ID获取多媒体
// @Description 根据ID获取家族多媒体文件信息
// @Tags 多媒体管理
// @Produce json
// @Param id query int true "多媒体ID"
// @Success 200 {object} mediaDomain.Media
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /media/getById [get]
func (h *MediaHandler) GetByID(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	m, err := h.appService.GetMediaByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, m)
}

// GetByPersonID 根据人员ID获取多媒体
// @Summary 根据人员ID获取多媒体
// @Description 根据人员ID获取其所有多媒体文件
// @Tags 多媒体管理
// @Produce json
// @Param personId query int true "人员ID"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /media/getByPersonId [get]
func (h *MediaHandler) GetByPersonID(c *gin.Context) {
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

	mediaItems, total, err := h.appService.GetMediaByPersonID(personID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  mediaItems,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetByEventID 根据事件ID获取多媒体
// @Summary 根据事件ID获取多媒体
// @Description 根据事件ID获取其所有多媒体文件
// @Tags 多媒体管理
// @Produce json
// @Param eventId query int true "事件ID"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /media/getByEventId [get]
func (h *MediaHandler) GetByEventID(c *gin.Context) {
	eventIDStr := c.Query("eventId")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
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

	mediaItems, total, err := h.appService.GetMediaByEventID(eventID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  mediaItems,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetByType 根据媒体类型获取多媒体
// @Summary 根据媒体类型获取多媒体
// @Description 根据多媒体类型获取所有多媒体文件
// @Tags 多媒体管理
// @Produce json
// @Param type query string true "多媒体类型"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /media/getByType [get]
func (h *MediaHandler) GetByType(c *gin.Context) {
	mediaType := c.Query("type")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	if mediaType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid media type"})
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

	mediaItems, total, err := h.appService.GetMediaByType(mediaType, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  mediaItems,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetPublic 获取公开的多媒体
// @Summary 获取公开的多媒体
// @Description 获取所有公开的家族多媒体文件
// @Tags 多媒体管理
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /media/getPublic [get]
func (h *MediaHandler) GetPublic(c *gin.Context) {
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

	mediaItems, total, err := h.appService.GetPublicMedia(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  mediaItems,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}
