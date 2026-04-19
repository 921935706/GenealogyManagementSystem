package event

import (
	"GenealogyManagementSystem/internal/domain/event"
)

// ApplicationService 事件应用服务
type ApplicationService struct {
	eventService *event.Service
}

// NewApplicationService 创建事件应用服务实例
func NewApplicationService(eventService *event.Service) *ApplicationService {
	return &ApplicationService{eventService: eventService}
}

// CreateEvent 创建事件
func (s *ApplicationService) CreateEvent(e *event.Event) error {
	return s.eventService.CreateEvent(e)
}

// UpdateEvent 更新事件
func (s *ApplicationService) UpdateEvent(e *event.Event) error {
	return s.eventService.UpdateEvent(e)
}

// DeleteEvent 删除事件
func (s *ApplicationService) DeleteEvent(id int) error {
	return s.eventService.DeleteEvent(id)
}

// GetEventByID 根据ID获取事件
func (s *ApplicationService) GetEventByID(id int) (*event.Event, error) {
	return s.eventService.GetEventByID(id)
}

// GetEventsByPersonID 根据人员ID获取事件
func (s *ApplicationService) GetEventsByPersonID(personID int, page, pageSize int) ([]*event.Event, int, error) {
	return s.eventService.GetEventsByPersonID(personID, page, pageSize)
}

// GetEventsByType 根据事件类型获取事件
func (s *ApplicationService) GetEventsByType(eventType string, page, pageSize int) ([]*event.Event, int, error) {
	return s.eventService.GetEventsByType(eventType, page, pageSize)
}

// GetImportantEvents 获取重要事件
func (s *ApplicationService) GetImportantEvents(page, pageSize int) ([]*event.Event, int, error) {
	return s.eventService.GetImportantEvents(page, pageSize)
}
