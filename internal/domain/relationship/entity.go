package relationship

import (
	"time"
)

// ParentChild 亲子关系实体
type ParentChild struct {
	ID           int       `json:"id"`
	ParentID     int       `json:"parent_id"`
	ChildID      int       `json:"child_id"`
	RelationType string    `json:"relation_type"`
	IsPrimary    bool      `json:"is_primary"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
}
