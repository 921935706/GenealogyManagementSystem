package family_relation

import (
	"time"
)

// FamilyRelation 家庭关系实体
type FamilyRelation struct {
	ID           int       `json:"id"`
	FamilyID     string    `json:"family_id"`
	PersonID     int       `json:"person_id"`
	SpouseID     int       `json:"spouse_id"`
	MarriageDate time.Time `json:"marriage_date"`
	MarriagePlace string    `json:"marriage_place"`
	DivorceDate  time.Time `json:"divorce_date"`
	RelationType string    `json:"relation_type"`
	IsPrimary    bool      `json:"is_primary"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
}
