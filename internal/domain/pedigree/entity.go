package pedigree

// Pedigree 世系实体
type Pedigree struct {
	ID          int    `json:"id"`
	PersonID    int    `json:"person_id"`
	FatherID    int    `json:"father_id"`
	MotherID    int    `json:"mother_id"`
	SpouseID    int    `json:"spouse_id"`
	Generation  int    `json:"generation"`
	LineagePath string `json:"lineage_path"`
	Depth       int    `json:"depth"`
}
