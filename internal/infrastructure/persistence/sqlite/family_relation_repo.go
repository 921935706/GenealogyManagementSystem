package sqlite

import (
	"GenealogyManagementSystem/internal/domain/family_relation"
	"database/sql"
)

// FamilyRelationRepository 家庭关系仓储实现
type FamilyRelationRepository struct {
	db *sql.DB
}

// NewFamilyRelationRepository 创建家庭关系仓储实例
func NewFamilyRelationRepository(db *sql.DB) *FamilyRelationRepository {
	return &FamilyRelationRepository{db: db}
}

// Create 创建家庭关系
func (r *FamilyRelationRepository) Create(relation *family_relation.FamilyRelation) error {
	query := `INSERT INTO family_relation (family_id, person_id, spouse_id, marriage_date, marriage_place, divorce_date, relation_type, is_primary, notes, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, relation.FamilyID, relation.PersonID, relation.SpouseID, relation.MarriageDate, relation.MarriagePlace, relation.DivorceDate, relation.RelationType, relation.IsPrimary, relation.Notes, relation.CreatedAt)
	return err
}

// Update 更新家庭关系
func (r *FamilyRelationRepository) Update(relation *family_relation.FamilyRelation) error {
	query := `UPDATE family_relation SET family_id = ?, person_id = ?, spouse_id = ?, marriage_date = ?, marriage_place = ?, divorce_date = ?, relation_type = ?, is_primary = ?, notes = ? WHERE id = ?`
	_, err := r.db.Exec(query, relation.FamilyID, relation.PersonID, relation.SpouseID, relation.MarriageDate, relation.MarriagePlace, relation.DivorceDate, relation.RelationType, relation.IsPrimary, relation.Notes, relation.ID)
	return err
}

// Delete 删除家庭关系
func (r *FamilyRelationRepository) Delete(id int) error {
	query := `DELETE FROM family_relation WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// GetByID 根据ID获取家庭关系
func (r *FamilyRelationRepository) GetByID(id int) (*family_relation.FamilyRelation, error) {
	query := `SELECT id, family_id, person_id, spouse_id, marriage_date, marriage_place, divorce_date, relation_type, is_primary, notes, created_at FROM family_relation WHERE id = ?`
	row := r.db.QueryRow(query, id)

	relation := &family_relation.FamilyRelation{}
	err := row.Scan(&relation.ID, &relation.FamilyID, &relation.PersonID, &relation.SpouseID, &relation.MarriageDate, &relation.MarriagePlace, &relation.DivorceDate, &relation.RelationType, &relation.IsPrimary, &relation.Notes, &relation.CreatedAt)
	if err != nil {
		return nil, err
	}

	return relation, nil
}

// GetByPersonID 根据人员ID获取家庭关系
func (r *FamilyRelationRepository) GetByPersonID(personID int) ([]*family_relation.FamilyRelation, error) {
	query := `SELECT id, family_id, person_id, spouse_id, marriage_date, marriage_place, divorce_date, relation_type, is_primary, notes, created_at FROM family_relation WHERE person_id = ?`
	rows, err := r.db.Query(query, personID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	relations := []*family_relation.FamilyRelation{}
	for rows.Next() {
		relation := &family_relation.FamilyRelation{}
		err := rows.Scan(&relation.ID, &relation.FamilyID, &relation.PersonID, &relation.SpouseID, &relation.MarriageDate, &relation.MarriagePlace, &relation.DivorceDate, &relation.RelationType, &relation.IsPrimary, &relation.Notes, &relation.CreatedAt)
		if err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}

	return relations, nil
}

// GetBySpouseID 根据配偶ID获取家庭关系
func (r *FamilyRelationRepository) GetBySpouseID(spouseID int) ([]*family_relation.FamilyRelation, error) {
	query := `SELECT id, family_id, person_id, spouse_id, marriage_date, marriage_place, divorce_date, relation_type, is_primary, notes, created_at FROM family_relation WHERE spouse_id = ?`
	rows, err := r.db.Query(query, spouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	relations := []*family_relation.FamilyRelation{}
	for rows.Next() {
		relation := &family_relation.FamilyRelation{}
		err := rows.Scan(&relation.ID, &relation.FamilyID, &relation.PersonID, &relation.SpouseID, &relation.MarriageDate, &relation.MarriagePlace, &relation.DivorceDate, &relation.RelationType, &relation.IsPrimary, &relation.Notes, &relation.CreatedAt)
		if err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}

	return relations, nil
}

// GetByFamilyID 根据家庭ID获取家庭关系
func (r *FamilyRelationRepository) GetByFamilyID(familyID string) ([]*family_relation.FamilyRelation, error) {
	query := `SELECT id, family_id, person_id, spouse_id, marriage_date, marriage_place, divorce_date, relation_type, is_primary, notes, created_at FROM family_relation WHERE family_id = ?`
	rows, err := r.db.Query(query, familyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	relations := []*family_relation.FamilyRelation{}
	for rows.Next() {
		relation := &family_relation.FamilyRelation{}
		err := rows.Scan(&relation.ID, &relation.FamilyID, &relation.PersonID, &relation.SpouseID, &relation.MarriageDate, &relation.MarriagePlace, &relation.DivorceDate, &relation.RelationType, &relation.IsPrimary, &relation.Notes, &relation.CreatedAt)
		if err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}

	return relations, nil
}
