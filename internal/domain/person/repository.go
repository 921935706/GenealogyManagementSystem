package person

// Repository 人员仓储接口
type Repository interface {
	Create(person *Person) error
	Update(person *Person) error
	Delete(id int) error
	GetByID(id int) (*Person, error)
	GetByUUID(uuid string) (*Person, error)
	List(page, pageSize int) ([]*Person, int, error)
	Search(name string, page, pageSize int) ([]*Person, int, error)
}
