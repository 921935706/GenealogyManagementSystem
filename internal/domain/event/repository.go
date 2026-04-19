package event

// Repository 事件仓储接口
type Repository interface {
	Create(event *Event) error
	Update(event *Event) error
	Delete(id int) error
	GetByID(id int) (*Event, error)
	GetByPersonID(personID int, page, pageSize int) ([]*Event, int, error)
	GetByType(eventType string, page, pageSize int) ([]*Event, int, error)
	GetImportantEvents(page, pageSize int) ([]*Event, int, error)
}
