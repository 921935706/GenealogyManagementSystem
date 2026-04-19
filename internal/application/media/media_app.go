package media

import (
	"GenealogyManagementSystem/internal/domain/media"
)

// ApplicationService 多媒体应用服务
type ApplicationService struct {
	mediaService *media.Service
}

// NewApplicationService 创建多媒体应用服务实例
func NewApplicationService(mediaService *media.Service) *ApplicationService {
	return &ApplicationService{mediaService: mediaService}
}

// CreateMedia 创建多媒体
func (s *ApplicationService) CreateMedia(m *media.Media) error {
	return s.mediaService.CreateMedia(m)
}

// UpdateMedia 更新多媒体
func (s *ApplicationService) UpdateMedia(m *media.Media) error {
	return s.mediaService.UpdateMedia(m)
}

// DeleteMedia 删除多媒体
func (s *ApplicationService) DeleteMedia(id int) error {
	return s.mediaService.DeleteMedia(id)
}

// GetMediaByID 根据ID获取多媒体
func (s *ApplicationService) GetMediaByID(id int) (*media.Media, error) {
	return s.mediaService.GetMediaByID(id)
}

// GetMediaByPersonID 根据人员ID获取多媒体
func (s *ApplicationService) GetMediaByPersonID(personID int, page, pageSize int) ([]*media.Media, int, error) {
	return s.mediaService.GetMediaByPersonID(personID, page, pageSize)
}

// GetMediaByEventID 根据事件ID获取多媒体
func (s *ApplicationService) GetMediaByEventID(eventID int, page, pageSize int) ([]*media.Media, int, error) {
	return s.mediaService.GetMediaByEventID(eventID, page, pageSize)
}

// GetMediaByType 根据媒体类型获取多媒体
func (s *ApplicationService) GetMediaByType(mediaType string, page, pageSize int) ([]*media.Media, int, error) {
	return s.mediaService.GetMediaByType(mediaType, page, pageSize)
}

// GetPublicMedia 获取公开的多媒体
func (s *ApplicationService) GetPublicMedia(page, pageSize int) ([]*media.Media, int, error) {
	return s.mediaService.GetPublicMedia(page, pageSize)
}
