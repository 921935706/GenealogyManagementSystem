package person

// Service 人员领域服务
type Service struct {
	repo Repository
}

// NewService 创建人员服务实例
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreatePerson 创建人员
func (s *Service) CreatePerson(person *Person) error {
	return s.repo.Create(person)
}

// UpdatePerson 更新人员信息
func (s *Service) UpdatePerson(person *Person) error {
	return s.repo.Update(person)
}

// DeletePerson 删除人员
func (s *Service) DeletePerson(id int) error {
	return s.repo.Delete(id)
}

// GetPersonByID 根据ID获取人员
func (s *Service) GetPersonByID(id int) (*Person, error) {
	return s.repo.GetByID(id)
}

// GetPersonByUUID 根据UUID获取人员
func (s *Service) GetPersonByUUID(uuid string) (*Person, error) {
	return s.repo.GetByUUID(uuid)
}

// ListPersons 列出人员
func (s *Service) ListPersons(page, pageSize int) ([]*Person, int, error) {
	return s.repo.List(page, pageSize)
}

// SearchPersons 搜索人员
func (s *Service) SearchPersons(name string, page, pageSize int) ([]*Person, int, error) {
	return s.repo.Search(name, page, pageSize)
}
