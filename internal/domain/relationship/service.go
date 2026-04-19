package relationship

// Service 关系领域服务
type Service struct {
	repo Repository
}

// NewService 创建关系服务实例
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateParentChild 创建亲子关系
func (s *Service) CreateParentChild(relation *ParentChild) error {
	return s.repo.CreateParentChild(relation)
}

// DeleteParentChild 删除亲子关系
func (s *Service) DeleteParentChild(parentID, childID int) error {
	return s.repo.DeleteParentChild(parentID, childID)
}

// GetParentChild 获取亲子关系
func (s *Service) GetParentChild(parentID, childID int) (*ParentChild, error) {
	return s.repo.GetParentChild(parentID, childID)
}

// GetChildrenByParentID 获取父母的子女
func (s *Service) GetChildrenByParentID(parentID int) ([]*ParentChild, error) {
	return s.repo.GetChildrenByParentID(parentID)
}

// GetParentsByChildID 获取子女的父母
func (s *Service) GetParentsByChildID(childID int) ([]*ParentChild, error) {
	return s.repo.GetParentsByChildID(childID)
}
