package handler

import (
	"GenealogyManagementSystem/internal/application/relationship"
	relationDomain "GenealogyManagementSystem/internal/domain/relationship"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RelationshipHandler 关系处理器
type RelationshipHandler struct {
	appService *relationship.ApplicationService
}

// NewRelationshipHandler 创建关系处理器实例
func NewRelationshipHandler(appService *relationship.ApplicationService) *RelationshipHandler {
	return &RelationshipHandler{appService: appService}
}

// CreateParentChild 创建亲子关系
// @Summary 创建亲子关系
// @Description 创建父母与子女之间的关系
// @Tags 亲子关系管理
// @Accept json
// @Produce json
// @Param relation body relationDomain.ParentChild true "亲子关系信息"
// @Success 201 {object} relationDomain.ParentChild
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /relationship/createParentChild [post]
func (h *RelationshipHandler) CreateParentChild(c *gin.Context) {
	var relation relationDomain.ParentChild
	if err := c.ShouldBindJSON(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.CreateParentChild(&relation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, relation)
}

// DeleteParentChild 删除亲子关系
// @Summary 删除亲子关系
// @Description 删除父母与子女之间的关系
// @Tags 亲子关系管理
// @Produce json
// @Param parentId query int true "父母ID"
// @Param childId query int true "子女ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /relationship/deleteParentChild [delete]
func (h *RelationshipHandler) DeleteParentChild(c *gin.Context) {
	parentIDStr := c.Query("parentId")
	childIDStr := c.Query("childId")

	parentID, err := strconv.Atoi(parentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent ID"})
		return
	}

	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid child ID"})
		return
	}

	if err := h.appService.DeleteParentChild(parentID, childID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetParentChild 获取亲子关系
// @Summary 获取亲子关系
// @Description 获取父母与子女之间的关系
// @Tags 亲子关系管理
// @Produce json
// @Param parentId query int true "父母ID"
// @Param childId query int true "子女ID"
// @Success 200 {object} relationDomain.ParentChild
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /relationship/getParentChild [get]
func (h *RelationshipHandler) GetParentChild(c *gin.Context) {
	parentIDStr := c.Query("parentId")
	childIDStr := c.Query("childId")

	parentID, err := strconv.Atoi(parentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent ID"})
		return
	}

	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid child ID"})
		return
	}

	relation, err := h.appService.GetParentChild(parentID, childID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relation)
}

// GetChildrenByParentID 获取父母的子女
// @Summary 获取父母的子女
// @Description 根据父母ID获取其所有子女
// @Tags 亲子关系管理
// @Produce json
// @Param parentId query int true "父母ID"
// @Success 200 {array} relationDomain.ParentChild
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /relationship/getChildrenByParentId [get]
func (h *RelationshipHandler) GetChildrenByParentID(c *gin.Context) {
	parentIDStr := c.Query("parentId")
	parentID, err := strconv.Atoi(parentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent ID"})
		return
	}

	relations, err := h.appService.GetChildrenByParentID(parentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relations)
}

// GetParentsByChildID 获取子女的父母
// @Summary 获取子女的父母
// @Description 根据子女ID获取其所有父母
// @Tags 亲子关系管理
// @Produce json
// @Param childId query int true "子女ID"
// @Success 200 {array} relationDomain.ParentChild
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /relationship/getParentsByChildId [get]
func (h *RelationshipHandler) GetParentsByChildID(c *gin.Context) {
	childIDStr := c.Query("childId")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid child ID"})
		return
	}

	relations, err := h.appService.GetParentsByChildID(childID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relations)
}
