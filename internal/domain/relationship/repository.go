package relationship

// Repository 关系仓储接口
type Repository interface {
	CreateParentChild(relation *ParentChild) error
	DeleteParentChild(parentID, childID int) error
	GetParentChild(parentID, childID int) (*ParentChild, error)
	GetChildrenByParentID(parentID int) ([]*ParentChild, error)
	GetParentsByChildID(childID int) ([]*ParentChild, error)
}
