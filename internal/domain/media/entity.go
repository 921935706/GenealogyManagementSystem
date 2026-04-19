package media

import (
	"time"
)

// Media 多媒体实体
type Media struct {
	ID            int       `json:"id"`
	PersonID      int       `json:"person_id"`
	EventID       int       `json:"event_id"`
	MediaType     string    `json:"media_type"`
	FileName      string    `json:"file_name"`
	FileURL       string    `json:"file_url"`
	FileSize      int       `json:"file_size"`
	MimeType      string    `json:"mime_type"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	UploaderID    int       `json:"uploader_id"`
	IsPublic      bool      `json:"is_public"`
	UploadedAt    time.Time `json:"uploaded_at"`
}
