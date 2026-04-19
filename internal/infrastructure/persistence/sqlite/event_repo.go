package sqlite

import (
	"GenealogyManagementSystem/internal/domain/event"
	"database/sql"
)

// EventRepository 事件仓储实现
type EventRepository struct {
	db *sql.DB
}

// NewEventRepository 创建事件仓储实例
func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

// Create 创建事件
func (r *EventRepository) Create(e *event.Event) error {
	query := `INSERT INTO event (person_id, event_type, event_name, event_date, event_lunar, event_place, event_place_longitude, event_place_latitude, description, is_important, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, e.PersonID, e.EventType, e.EventName, e.EventDate, e.EventLunar, e.EventPlace, e.EventPlaceLongitude, e.EventPlaceLatitude, e.Description, e.IsImportant, e.CreatedAt)
	return err
}

// Update 更新事件
func (r *EventRepository) Update(e *event.Event) error {
	query := `UPDATE event SET person_id = ?, event_type = ?, event_name = ?, event_date = ?, event_lunar = ?, event_place = ?, event_place_longitude = ?, event_place_latitude = ?, description = ?, is_important = ? WHERE id = ?`
	_, err := r.db.Exec(query, e.PersonID, e.EventType, e.EventName, e.EventDate, e.EventLunar, e.EventPlace, e.EventPlaceLongitude, e.EventPlaceLatitude, e.Description, e.IsImportant, e.ID)
	return err
}

// Delete 删除事件
func (r *EventRepository) Delete(id int) error {
	query := `DELETE FROM event WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// GetByID 根据ID获取事件
func (r *EventRepository) GetByID(id int) (*event.Event, error) {
	query := `SELECT id, person_id, event_type, event_name, event_date, event_lunar, event_place, event_place_longitude, event_place_latitude, description, is_important, created_at FROM event WHERE id = ?`
	row := r.db.QueryRow(query, id)

	e := &event.Event{}
	err := row.Scan(&e.ID, &e.PersonID, &e.EventType, &e.EventName, &e.EventDate, &e.EventLunar, &e.EventPlace, &e.EventPlaceLongitude, &e.EventPlaceLatitude, &e.Description, &e.IsImportant, &e.CreatedAt)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// GetByPersonID 根据人员ID获取事件
func (r *EventRepository) GetByPersonID(personID int, page, pageSize int) ([]*event.Event, int, error) {
	offset := (page - 1) * pageSize

	// 获取总数
	countQuery := `SELECT COUNT(*) FROM event WHERE person_id = ?`
	var total int
	err := r.db.QueryRow(countQuery, personID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	query := `SELECT id, person_id, event_type, event_name, event_date, event_lunar, event_place, event_place_longitude, event_place_latitude, description, is_important, created_at FROM event WHERE person_id = ? LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, personID, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	events := []*event.Event{}
	for rows.Next() {
		e := &event.Event{}
		err := rows.Scan(&e.ID, &e.PersonID, &e.EventType, &e.EventName, &e.EventDate, &e.EventLunar, &e.EventPlace, &e.EventPlaceLongitude, &e.EventPlaceLatitude, &e.Description, &e.IsImportant, &e.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		events = append(events, e)
	}

	return events, total, nil
}

// GetByType 根据事件类型获取事件
func (r *EventRepository) GetByType(eventType string, page, pageSize int) ([]*event.Event, int, error) {
	offset := (page - 1) * pageSize

	// 获取总数
	countQuery := `SELECT COUNT(*) FROM event WHERE event_type = ?`
	var total int
	err := r.db.QueryRow(countQuery, eventType).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	query := `SELECT id, person_id, event_type, event_name, event_date, event_lunar, event_place, event_place_longitude, event_place_latitude, description, is_important, created_at FROM event WHERE event_type = ? LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, eventType, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	events := []*event.Event{}
	for rows.Next() {
		e := &event.Event{}
		err := rows.Scan(&e.ID, &e.PersonID, &e.EventType, &e.EventName, &e.EventDate, &e.EventLunar, &e.EventPlace, &e.EventPlaceLongitude, &e.EventPlaceLatitude, &e.Description, &e.IsImportant, &e.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		events = append(events, e)
	}

	return events, total, nil
}

// GetImportantEvents 获取重要事件
func (r *EventRepository) GetImportantEvents(page, pageSize int) ([]*event.Event, int, error) {
	offset := (page - 1) * pageSize

	// 获取总数
	countQuery := `SELECT COUNT(*) FROM event WHERE is_important = true`
	var total int
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	query := `SELECT id, person_id, event_type, event_name, event_date, event_lunar, event_place, event_place_longitude, event_place_latitude, description, is_important, created_at FROM event WHERE is_important = true LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	events := []*event.Event{}
	for rows.Next() {
		e := &event.Event{}
		err := rows.Scan(&e.ID, &e.PersonID, &e.EventType, &e.EventName, &e.EventDate, &e.EventLunar, &e.EventPlace, &e.EventPlaceLongitude, &e.EventPlaceLatitude, &e.Description, &e.IsImportant, &e.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		events = append(events, e)
	}

	return events, total, nil
}
