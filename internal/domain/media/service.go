package media

// Service 多媒体领域服务
type Service struct {
	repo Repository
}

// NewService 创建多媒体服务实例
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateMedia 创建多媒体
func (s *Service) CreateMedia(media *Media) error {
	return s.repo.Create(media)
}

// UpdateMedia 更新多媒体
func (s *Service) UpdateMedia(media *Media) error {
	return s.repo.Update(media)
}

// DeleteMedia 删除多媒体
func (s *Service) DeleteMedia(id int) error {
	return s.repo.Delete(id)
}

// GetMediaByID 根据ID获取多媒体
func (s *Service) GetMediaByID(id int) (*Media, error) {
	return s.repo.GetByID(id)
}

// GetMediaByPersonID 根据人员ID获取多媒体
func (s *Service) GetMediaByPersonID(personID int, page, pageSize int) ([]*Media, int, error) {
	return s.repo.GetByPersonID(personID, page, pageSize)
}

// GetMediaByEventID 根据事件ID获取多媒体
func (s *Service) GetMediaByEventID(eventID int, page, pageSize int) ([]*Media, int, error) {
	return s.repo.GetByEventID(eventID, page, pageSize)
}

// GetMediaByType 根据媒体类型获取多媒体
func (s *Service) GetMediaByType(mediaType string, page, pageSize int) ([]*Media, int, error) {
	return s.repo.GetByType(mediaType, page, pageSize)
}

// GetPublicMedia 获取公开的多媒体
func (s *Service) GetPublicMedia(page, pageSize int) ([]*Media, int, error) {
	return s.repo.GetPublicMedia(page, pageSize)
}
