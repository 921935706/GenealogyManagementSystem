package event

import (
	"time"
)

// Event 事件实体
type Event struct {
	ID                  int       `json:"id"`
	PersonID            int       `json:"person_id"`
	EventType           string    `json:"event_type"`
	EventName           string    `json:"event_name"`
	EventDate           time.Time `json:"event_date"`
	EventLunar          string    `json:"event_lunar"`
	EventPlace          string    `json:"event_place"`
	EventPlaceLongitude float64   `json:"event_place_longitude"`
	EventPlaceLatitude  float64   `json:"event_place_latitude"`
	Description         string    `json:"description"`
	IsImportant         bool      `json:"is_important"`
	CreatedAt           time.Time `json:"created_at"`
}
