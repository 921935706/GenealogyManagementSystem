package family_relation

// Repository 家庭关系仓储接口
type Repository interface {
	Create(relation *FamilyRelation) error
	Update(relation *FamilyRelation) error
	Delete(id int) error
	GetByID(id int) (*FamilyRelation, error)
	GetByPersonID(personID int) ([]*FamilyRelation, error)
	GetBySpouseID(spouseID int) ([]*FamilyRelation, error)
	GetByFamilyID(familyID string) ([]*FamilyRelation, error)
}
