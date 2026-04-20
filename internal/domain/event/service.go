package event

// Service 事件领域服务
type Service struct {
	repo Repository
}

// NewService 创建事件服务实例
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateEvent 创建事件
func (s *Service) CreateEvent(event *Event) error {
	return s.repo.Create(event)
}

// UpdateEvent 更新事件
func (s *Service) UpdateEvent(event *Event) error {
	return s.repo.Update(event)
}

// DeleteEvent 删除事件
func (s *Service) DeleteEvent(id int) error {
	return s.repo.Delete(id)
}

// GetEventByID 根据ID获取事件
func (s *Service) GetEventByID(id int) (*Event, error) {
	return s.repo.GetByID(id)
}

// GetEventsByPersonID 根据人员ID获取事件
func (s *Service) GetEventsByPersonID(personID int, page, pageSize int) ([]*Event, int, error) {
	return s.repo.GetByPersonID(personID, page, pageSize)
}

// GetEventsByType 根据事件类型获取事件
func (s *Service) GetEventsByType(eventType string, page, pageSize int) ([]*Event, int, error) {
	return s.repo.GetByType(eventType, page, pageSize)
}

// GetImportantEvents 获取重要事件
func (s *Service) GetImportantEvents(page, pageSize int) ([]*Event, int, error) {
	return s.repo.GetImportantEvents(page, pageSize)
}
