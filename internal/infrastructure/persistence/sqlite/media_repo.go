package sqlite

import (
	"GenealogyManagementSystem/internal/domain/media"
	"database/sql"
)

// MediaRepository 多媒体仓储实现
type MediaRepository struct {
	db *sql.DB
}

// NewMediaRepository 创建多媒体仓储实例
func NewMediaRepository(db *sql.DB) *MediaRepository {
	return &MediaRepository{db: db}
}

// Create 创建多媒体
func (r *MediaRepository) Create(m *media.Media) error {
	query := `INSERT INTO media (person_id, event_id, media_type, file_name, file_url, file_size, mime_type, thumbnail_url, title, description, uploader_id, is_public, uploaded_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, m.PersonID, m.EventID, m.MediaType, m.FileName, m.FileURL, m.FileSize, m.MimeType, m.ThumbnailURL, m.Title, m.Description, m.UploaderID, m.IsPublic, m.UploadedAt)
	return err
}

// Update 更新多媒体
func (r *MediaRepository) Update(m *media.Media) error {
	query := `UPDATE media SET person_id = ?, event_id = ?, media_type = ?, file_name = ?, file_url = ?, file_size = ?, mime_type = ?, thumbnail_url = ?, title = ?, description = ?, uploader_id = ?, is_public = ? WHERE id = ?`
	_, err := r.db.Exec(query, m.PersonID, m.EventID, m.MediaType, m.FileName, m.FileURL, m.FileSize, m.MimeType, m.ThumbnailURL, m.Title, m.Description, m.UploaderID, m.IsPublic, m.ID)
	return err
}

// Delete 删除多媒体
func (r *MediaRepository) Delete(id int) error {
	query := `DELETE FROM media WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// GetByID 根据ID获取多媒体
func (r *MediaRepository) GetByID(id int) (*media.Media, error) {
	query := `SELECT id, person_id, event_id, media_type, file_name, file_url, file_size, mime_type, thumbnail_url, title, description, uploader_id, is_public, uploaded_at FROM media WHERE id = ?`
	row := r.db.QueryRow(query, id)

	m := &media.Media{}
	err := row.Scan(&m.ID, &m.PersonID, &m.EventID, &m.MediaType, &m.FileName, &m.FileURL, &m.FileSize, &m.MimeType, &m.ThumbnailURL, &m.Title, &m.Description, &m.UploaderID, &m.IsPublic, &m.UploadedAt)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// GetByPersonID 根据人员ID获取多媒体
func (r *MediaRepository) GetByPersonID(personID int, page, pageSize int) ([]*media.Media, int, error) {
	offset := (page - 1) * pageSize

	// 获取总数
	countQuery := `SELECT COUNT(*) FROM media WHERE person_id = ?`
	var total int
	err := r.db.QueryRow(countQuery, personID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	query := `SELECT id, person_id, event_id, media_type, file_name, file_url, file_size, mime_type, thumbnail_url, title, description, uploader_id, is_public, uploaded_at FROM media WHERE person_id = ? LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, personID, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	mediaItems := []*media.Media{}
	for rows.Next() {
		m := &media.Media{}
		err := rows.Scan(&m.ID, &m.PersonID, &m.EventID, &m.MediaType, &m.FileName, &m.FileURL, &m.FileSize, &m.MimeType, &m.ThumbnailURL, &m.Title, &m.Description, &m.UploaderID, &m.IsPublic, &m.UploadedAt)
		if err != nil {
			return nil, 0, err
		}
		mediaItems = append(mediaItems, m)
	}

	return mediaItems, total, nil
}

// GetByEventID 根据事件ID获取多媒体
func (r *MediaRepository) GetByEventID(eventID int, page, pageSize int) ([]*media.Media, int, error) {
	offset := (page - 1) * pageSize

	// 获取总数
	countQuery := `SELECT COUNT(*) FROM media WHERE event_id = ?`
	var total int
	err := r.db.QueryRow(countQuery, eventID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	query := `SELECT id, person_id, event_id, media_type, file_name, file_url, file_size, mime_type, thumbnail_url, title, description, uploader_id, is_public, uploaded_at FROM media WHERE event_id = ? LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, eventID, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	mediaItems := []*media.Media{}
	for rows.Next() {
		m := &media.Media{}
		err := rows.Scan(&m.ID, &m.PersonID, &m.EventID, &m.MediaType, &m.FileName, &m.FileURL, &m.FileSize, &m.MimeType, &m.ThumbnailURL, &m.Title, &m.Description, &m.UploaderID, &m.IsPublic, &m.UploadedAt)
		if err != nil {
			return nil, 0, err
		}
		mediaItems = append(mediaItems, m)
	}

	return mediaItems, total, nil
}

// GetByType 根据媒体类型获取多媒体
func (r *MediaRepository) GetByType(mediaType string, page, pageSize int) ([]*media.Media, int, error) {
	offset := (page - 1) * pageSize

	// 获取总数
	countQuery := `SELECT COUNT(*) FROM media WHERE media_type = ?`
	var total int
	err := r.db.QueryRow(countQuery, mediaType).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	query := `SELECT id, person_id, event_id, media_type, file_name, file_url, file_size, mime_type, thumbnail_url, title, description, uploader_id, is_public, uploaded_at FROM media WHERE media_type = ? LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, mediaType, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	mediaItems := []*media.Media{}
	for rows.Next() {
		m := &media.Media{}
		err := rows.Scan(&m.ID, &m.PersonID, &m.EventID, &m.MediaType, &m.FileName, &m.FileURL, &m.FileSize, &m.MimeType, &m.ThumbnailURL, &m.Title, &m.Description, &m.UploaderID, &m.IsPublic, &m.UploadedAt)
		if err != nil {
			return nil, 0, err
		}
		mediaItems = append(mediaItems, m)
	}

	return mediaItems, total, nil
}

// GetPublicMedia 获取公开的多媒体
func (r *MediaRepository) GetPublicMedia(page, pageSize int) ([]*media.Media, int, error) {
	offset := (page - 1) * pageSize

	// 获取总数
	countQuery := `SELECT COUNT(*) FROM media WHERE is_public = true`
	var total int
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	query := `SELECT id, person_id, event_id, media_type, file_name, file_url, file_size, mime_type, thumbnail_url, title, description, uploader_id, is_public, uploaded_at FROM media WHERE is_public = true LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	mediaItems := []*media.Media{}
	for rows.Next() {
		m := &media.Media{}
		err := rows.Scan(&m.ID, &m.PersonID, &m.EventID, &m.MediaType, &m.FileName, &m.FileURL, &m.FileSize, &m.MimeType, &m.ThumbnailURL, &m.Title, &m.Description, &m.UploaderID, &m.IsPublic, &m.UploadedAt)
		if err != nil {
			return nil, 0, err
		}
		mediaItems = append(mediaItems, m)
	}

	return mediaItems, total, nil
}
