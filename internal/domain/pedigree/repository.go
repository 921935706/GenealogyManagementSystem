package pedigree

// Repository 世系仓储接口
type Repository interface {
	Create(pedigree *Pedigree) error
	Update(pedigree *Pedigree) error
	Delete(id int) error
	GetByPersonID(personID int) (*Pedigree, error)
	GetByID(id int) (*Pedigree, error)
	GetByGeneration(generation int) ([]*Pedigree, error)
	GetByLineagePath(path string) ([]*Pedigree, error)
}
