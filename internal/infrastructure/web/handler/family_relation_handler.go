package handler

import (
	"GenealogyManagementSystem/internal/application/family_relation"
	familyDomain "GenealogyManagementSystem/internal/domain/family_relation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FamilyRelationHandler 家庭关系处理器
type FamilyRelationHandler struct {
	appService *family_relation.ApplicationService
}

// NewFamilyRelationHandler 创建家庭关系处理器实例
func NewFamilyRelationHandler(appService *family_relation.ApplicationService) *FamilyRelationHandler {
	return &FamilyRelationHandler{appService: appService}
}

// Create 创建家庭关系
// @Summary 创建家庭关系
// @Description 创建夫妻/伴侣关系
// @Tags 家庭关系管理
// @Accept json
// @Produce json
// @Param relation body familyDomain.FamilyRelation true "家庭关系信息"
// @Success 201 {object} familyDomain.FamilyRelation
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /familyRelation/create [post]
func (h *FamilyRelationHandler) Create(c *gin.Context) {
	var relation familyDomain.FamilyRelation
	if err := c.ShouldBindJSON(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.CreateFamilyRelation(&relation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, relation)
}

// Update 更新家庭关系
// @Summary 更新家庭关系
// @Description 更新夫妻/伴侣关系信息
// @Tags 家庭关系管理
// @Accept json
// @Produce json
// @Param relation body familyDomain.FamilyRelation true "家庭关系信息"
// @Success 200 {object} familyDomain.FamilyRelation
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /familyRelation/update [put]
func (h *FamilyRelationHandler) Update(c *gin.Context) {
	var relation familyDomain.FamilyRelation
	if err := c.ShouldBindJSON(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.UpdateFamilyRelation(&relation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relation)
}

// Delete 删除家庭关系
// @Summary 删除家庭关系
// @Description 删除夫妻/伴侣关系
// @Tags 家庭关系管理
// @Produce json
// @Param id query int true "家庭关系ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /familyRelation/delete [delete]
func (h *FamilyRelationHandler) Delete(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.appService.DeleteFamilyRelation(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID 根据ID获取家庭关系
// @Summary 根据ID获取家庭关系
// @Description 根据ID获取夫妻/伴侣关系信息
// @Tags 家庭关系管理
// @Produce json
// @Param id query int true "家庭关系ID"
// @Success 200 {object} familyDomain.FamilyRelation
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /familyRelation/getById [get]
func (h *FamilyRelationHandler) GetByID(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	relation, err := h.appService.GetFamilyRelationByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relation)
}

// GetByPersonID 根据人员ID获取家庭关系
// @Summary 根据人员ID获取家庭关系
// @Description 根据人员ID获取其所有夫妻/伴侣关系
// @Tags 家庭关系管理
// @Produce json
// @Param personId query int true "人员ID"
// @Success 200 {array} familyDomain.FamilyRelation
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /familyRelation/getByPersonId [get]
func (h *FamilyRelationHandler) GetByPersonID(c *gin.Context) {
	personIDStr := c.Query("personId")
	personID, err := strconv.Atoi(personIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	relations, err := h.appService.GetFamilyRelationsByPersonID(personID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relations)
}

// GetBySpouseID 根据配偶ID获取家庭关系
// @Summary 根据配偶ID获取家庭关系
// @Description 根据配偶ID获取其所有夫妻/伴侣关系
// @Tags 家庭关系管理
// @Produce json
// @Param spouseId query int true "配偶ID"
// @Success 200 {array} familyDomain.FamilyRelation
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /familyRelation/getBySpouseId [get]
func (h *FamilyRelationHandler) GetBySpouseID(c *gin.Context) {
	spouseIDStr := c.Query("spouseId")
	spouseID, err := strconv.Atoi(spouseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spouse ID"})
		return
	}

	relations, err := h.appService.GetFamilyRelationsBySpouseID(spouseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relations)
}

// GetByFamilyID 根据家庭ID获取家庭关系
// @Summary 根据家庭ID获取家庭关系
// @Description 根据家庭ID获取其所有夫妻/伴侣关系
// @Tags 家庭关系管理
// @Produce json
// @Param familyId query string true "家庭ID"
// @Success 200 {array} familyDomain.FamilyRelation
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /familyRelation/getByFamilyId [get]
func (h *FamilyRelationHandler) GetByFamilyID(c *gin.Context) {
	familyID := c.Query("familyId")
	if familyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid family ID"})
		return
	}

	relations, err := h.appService.GetFamilyRelationsByFamilyID(familyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relations)
}
