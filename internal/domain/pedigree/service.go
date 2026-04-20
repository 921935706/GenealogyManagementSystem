package pedigree

// Service 世系领域服务
type Service struct {
	repo Repository
}

// NewService 创建世系服务实例
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreatePedigree 创建世系
func (s *Service) CreatePedigree(pedigree *Pedigree) error {
	return s.repo.Create(pedigree)
}

// UpdatePedigree 更新世系
func (s *Service) UpdatePedigree(pedigree *Pedigree) error {
	return s.repo.Update(pedigree)
}

// DeletePedigree 删除世系
func (s *Service) DeletePedigree(id int) error {
	return s.repo.Delete(id)
}

// GetPedigreeByPersonID 根据人员ID获取世系
func (s *Service) GetPedigreeByPersonID(personID int) (*Pedigree, error) {
	return s.repo.GetByPersonID(personID)
}

// GetPedigreeByID 根据ID获取世系
func (s *Service) GetPedigreeByID(id int) (*Pedigree, error) {
	return s.repo.GetByID(id)
}

// GetPedigreesByGeneration 根据世代获取世系
func (s *Service) GetPedigreesByGeneration(generation int) ([]*Pedigree, error) {
	return s.repo.GetByGeneration(generation)
}

// GetPedigreesByLineagePath 根据世系路径获取世系
func (s *Service) GetPedigreesByLineagePath(path string) ([]*Pedigree, error) {
	return s.repo.GetByLineagePath(path)
}
