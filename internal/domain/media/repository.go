package media

// Repository 多媒体仓储接口
type Repository interface {
	Create(media *Media) error
	Update(media *Media) error
	Delete(id int) error
	GetByID(id int) (*Media, error)
	GetByPersonID(personID int, page, pageSize int) ([]*Media, int, error)
	GetByEventID(eventID int, page, pageSize int) ([]*Media, int, error)
	GetByType(mediaType string, page, pageSize int) ([]*Media, int, error)
	GetPublicMedia(page, pageSize int) ([]*Media, int, error)
}
