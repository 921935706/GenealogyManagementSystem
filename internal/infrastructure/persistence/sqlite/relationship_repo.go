package sqlite

import (
	"GenealogyManagementSystem/internal/domain/relationship"
	"database/sql"
)

// RelationshipRepository 关系仓储实现
type RelationshipRepository struct {
	db *sql.DB
}

// NewRelationshipRepository 创建关系仓储实例
func NewRelationshipRepository(db *sql.DB) *RelationshipRepository {
	return &RelationshipRepository{db: db}
}

// CreateParentChild 创建亲子关系
func (r *RelationshipRepository) CreateParentChild(relation *relationship.ParentChild) error {
	query := `INSERT INTO parent_child (parent_id, child_id, relation_type, is_primary, notes, created_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, relation.ParentID, relation.ChildID, relation.RelationType, relation.IsPrimary, relation.Notes, relation.CreatedAt)
	return err
}

// DeleteParentChild 删除亲子关系
func (r *RelationshipRepository) DeleteParentChild(parentID, childID int) error {
	query := `DELETE FROM parent_child WHERE parent_id = ? AND child_id = ?`
	_, err := r.db.Exec(query, parentID, childID)
	return err
}

// GetParentChild 获取亲子关系
func (r *RelationshipRepository) GetParentChild(parentID, childID int) (*relationship.ParentChild, error) {
	query := `SELECT id, parent_id, child_id, relation_type, is_primary, notes, created_at FROM parent_child WHERE parent_id = ? AND child_id = ?`
	row := r.db.QueryRow(query, parentID, childID)

	relation := &relationship.ParentChild{}
	err := row.Scan(&relation.ID, &relation.ParentID, &relation.ChildID, &relation.RelationType, &relation.IsPrimary, &relation.Notes, &relation.CreatedAt)
	if err != nil {
		return nil, err
	}

	return relation, nil
}

// GetChildrenByParentID 获取父母的子女
func (r *RelationshipRepository) GetChildrenByParentID(parentID int) ([]*relationship.ParentChild, error) {
	query := `SELECT id, parent_id, child_id, relation_type, is_primary, notes, created_at FROM parent_child WHERE parent_id = ?`
	rows, err := r.db.Query(query, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	relations := []*relationship.ParentChild{}
	for rows.Next() {
		relation := &relationship.ParentChild{}
		err := rows.Scan(&relation.ID, &relation.ParentID, &relation.ChildID, &relation.RelationType, &relation.IsPrimary, &relation.Notes, &relation.CreatedAt)
		if err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}

	return relations, nil
}

// GetParentsByChildID 获取子女的父母
func (r *RelationshipRepository) GetParentsByChildID(childID int) ([]*relationship.ParentChild, error) {
	query := `SELECT id, parent_id, child_id, relation_type, is_primary, notes, created_at FROM parent_child WHERE child_id = ?`
	rows, err := r.db.Query(query, childID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	relations := []*relationship.ParentChild{}
	for rows.Next() {
		relation := &relationship.ParentChild{}
		err := rows.Scan(&relation.ID, &relation.ParentID, &relation.ChildID, &relation.RelationType, &relation.IsPrimary, &relation.Notes, &relation.CreatedAt)
		if err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}

	return relations, nil
}
