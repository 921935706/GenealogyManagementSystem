package person

import (
	"time"
)

// Person 人员实体
type Person struct {
	ID                   int       `json:"id"`
	UUID                 string    `json:"uuid"`
	Name                 string    `json:"name"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	GenerationName       string    `json:"generation_name"`
	Alias                string    `json:"alias"`
	EnglishName          string    `json:"english_name"`
	Gender               string    `json:"gender"`
	IsAlive              bool      `json:"is_alive"`
	BirthDate            time.Time `json:"birth_date"`
	BirthLunar           string    `json:"birth_lunar"`
	BirthPlace           string    `json:"birth_place"`
	BirthPlaceLongitude  float64   `json:"birth_place_longitude"`
	BirthPlaceLatitude   float64   `json:"birth_place_latitude"`
	DeathDate            time.Time `json:"death_date"`
	DeathLunar           string    `json:"death_lunar"`
	DeathPlace           string    `json:"death_place"`
	DeathPlaceLongitude  float64   `json:"death_place_longitude"`
	DeathPlaceLatitude   float64   `json:"death_place_latitude"`
	DeathCause           string    `json:"death_cause"`
	Generation           int       `json:"generation"`
	Occupation           string    `json:"occupation"`
	Education            string    `json:"education"`
	Biography            string    `json:"biography"`
	AvatarURL            string    `json:"avatar_url"`
	CoverPhotoURL        string    `json:"cover_photo_url"`
	IsPublic             bool      `json:"is_public"`
	VerificationStatus   string    `json:"verification_status"`
	CreatedBy            int       `json:"created_by"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
