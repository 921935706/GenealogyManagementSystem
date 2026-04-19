package family_relation

// Service 家庭关系领域服务
type Service struct {
	repo Repository
}

// NewService 创建家庭关系服务实例
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateFamilyRelation 创建家庭关系
func (s *Service) CreateFamilyRelation(relation *FamilyRelation) error {
	return s.repo.Create(relation)
}

// UpdateFamilyRelation 更新家庭关系
func (s *Service) UpdateFamilyRelation(relation *FamilyRelation) error {
	return s.repo.Update(relation)
}

// DeleteFamilyRelation 删除家庭关系
func (s *Service) DeleteFamilyRelation(id int) error {
	return s.repo.Delete(id)
}

// GetFamilyRelationByID 根据ID获取家庭关系
func (s *Service) GetFamilyRelationByID(id int) (*FamilyRelation, error) {
	return s.repo.GetByID(id)
}

// GetFamilyRelationsByPersonID 根据人员ID获取家庭关系
func (s *Service) GetFamilyRelationsByPersonID(personID int) ([]*FamilyRelation, error) {
	return s.repo.GetByPersonID(personID)
}

// GetFamilyRelationsBySpouseID 根据配偶ID获取家庭关系
func (s *Service) GetFamilyRelationsBySpouseID(spouseID int) ([]*FamilyRelation, error) {
	return s.repo.GetBySpouseID(spouseID)
}

// GetFamilyRelationsByFamilyID 根据家庭ID获取家庭关系
func (s *Service) GetFamilyRelationsByFamilyID(familyID string) ([]*FamilyRelation, error) {
	return s.repo.GetByFamilyID(familyID)
}
