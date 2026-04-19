package family_relation

import (
	"GenealogyManagementSystem/internal/domain/family_relation"
)

// ApplicationService 家庭关系应用服务
type ApplicationService struct {
	familyRelationService *family_relation.Service
}

// NewApplicationService 创建家庭关系应用服务实例
func NewApplicationService(familyRelationService *family_relation.Service) *ApplicationService {
	return &ApplicationService{familyRelationService: familyRelationService}
}

// CreateFamilyRelation 创建家庭关系
func (s *ApplicationService) CreateFamilyRelation(relation *family_relation.FamilyRelation) error {
	return s.familyRelationService.CreateFamilyRelation(relation)
}

// UpdateFamilyRelation 更新家庭关系
func (s *ApplicationService) UpdateFamilyRelation(relation *family_relation.FamilyRelation) error {
	return s.familyRelationService.UpdateFamilyRelation(relation)
}

// DeleteFamilyRelation 删除家庭关系
func (s *ApplicationService) DeleteFamilyRelation(id int) error {
	return s.familyRelationService.DeleteFamilyRelation(id)
}

// GetFamilyRelationByID 根据ID获取家庭关系
func (s *ApplicationService) GetFamilyRelationByID(id int) (*family_relation.FamilyRelation, error) {
	return s.familyRelationService.GetFamilyRelationByID(id)
}

// GetFamilyRelationsByPersonID 根据人员ID获取家庭关系
func (s *ApplicationService) GetFamilyRelationsByPersonID(personID int) ([]*family_relation.FamilyRelation, error) {
	return s.familyRelationService.GetFamilyRelationsByPersonID(personID)
}

// GetFamilyRelationsBySpouseID 根据配偶ID获取家庭关系
func (s *ApplicationService) GetFamilyRelationsBySpouseID(spouseID int) ([]*family_relation.FamilyRelation, error) {
	return s.familyRelationService.GetFamilyRelationsBySpouseID(spouseID)
}

// GetFamilyRelationsByFamilyID 根据家庭ID获取家庭关系
func (s *ApplicationService) GetFamilyRelationsByFamilyID(familyID string) ([]*family_relation.FamilyRelation, error) {
	return s.familyRelationService.GetFamilyRelationsByFamilyID(familyID)
}
