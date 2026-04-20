package relationship

import (
	"GenealogyManagementSystem/internal/domain/relationship"
)

// ApplicationService 关系应用服务
type ApplicationService struct {
	relationshipService *relationship.Service
}

// NewApplicationService 创建关系应用服务实例
func NewApplicationService(relationshipService *relationship.Service) *ApplicationService {
	return &ApplicationService{relationshipService: relationshipService}
}

// CreateParentChild 创建亲子关系
func (s *ApplicationService) CreateParentChild(relation *relationship.ParentChild) error {
	return s.relationshipService.CreateParentChild(relation)
}

// DeleteParentChild 删除亲子关系
func (s *ApplicationService) DeleteParentChild(parentID, childID int) error {
	return s.relationshipService.DeleteParentChild(parentID, childID)
}

// GetParentChild 获取亲子关系
func (s *ApplicationService) GetParentChild(parentID, childID int) (*relationship.ParentChild, error) {
	return s.relationshipService.GetParentChild(parentID, childID)
}

// GetChildrenByParentID 获取父母的子女
func (s *ApplicationService) GetChildrenByParentID(parentID int) ([]*relationship.ParentChild, error) {
	return s.relationshipService.GetChildrenByParentID(parentID)
}

// GetParentsByChildID 获取子女的父母
func (s *ApplicationService) GetParentsByChildID(childID int) ([]*relationship.ParentChild, error) {
	return s.relationshipService.GetParentsByChildID(childID)
}
