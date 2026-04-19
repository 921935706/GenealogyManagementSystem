package person

import (
	"GenealogyManagementSystem/internal/domain/person"
)

// ApplicationService 人员应用服务
type ApplicationService struct {
	personService *person.Service
}

// NewApplicationService 创建人员应用服务实例
func NewApplicationService(personService *person.Service) *ApplicationService {
	return &ApplicationService{personService: personService}
}

// CreatePerson 创建人员
func (s *ApplicationService) CreatePerson(p *person.Person) error {
	return s.personService.CreatePerson(p)
}

// UpdatePerson 更新人员信息
func (s *ApplicationService) UpdatePerson(p *person.Person) error {
	return s.personService.UpdatePerson(p)
}

// DeletePerson 删除人员
func (s *ApplicationService) DeletePerson(id int) error {
	return s.personService.DeletePerson(id)
}

// GetPersonByID 根据ID获取人员
func (s *ApplicationService) GetPersonByID(id int) (*person.Person, error) {
	return s.personService.GetPersonByID(id)
}

// GetPersonByUUID 根据UUID获取人员
func (s *ApplicationService) GetPersonByUUID(uuid string) (*person.Person, error) {
	return s.personService.GetPersonByUUID(uuid)
}

// ListPersons 列出人员
func (s *ApplicationService) ListPersons(page, pageSize int) ([]*person.Person, int, error) {
	return s.personService.ListPersons(page, pageSize)
}

// SearchPersons 搜索人员
func (s *ApplicationService) SearchPersons(name string, page, pageSize int) ([]*person.Person, int, error) {
	return s.personService.SearchPersons(name, page, pageSize)
}
