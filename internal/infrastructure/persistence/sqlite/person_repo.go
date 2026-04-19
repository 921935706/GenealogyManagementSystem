package sqlite

import (
	"GenealogyManagementSystem/internal/domain/person"
	"database/sql"
	"time"
)

// PersonRepository 人员仓储实现
type PersonRepository struct {
	db *sql.DB
}

// NewPersonRepository 创建人员仓储实例
func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

// Create 创建人员
func (r *PersonRepository) Create(p *person.Person) error {
	builder := NewInsertBuilder("person")
	query, args := builder.
		Columns(
			"uuid", "name", "first_name", "last_name", "generation_name", "alias", 
			"english_name", "gender", "is_alive", "birth_date", "birth_lunar", 
			"birth_place", "birth_place_longitude", "birth_place_latitude", 
			"death_date", "death_lunar", "death_place", "death_place_longitude", 
			"death_place_latitude", "death_cause", "generation", "occupation", 
			"education", "biography", "avatar_url", "cover_photo_url", "is_public", 
			"verification_status", "created_by", "created_at", "updated_at",
		).
		Values(
			p.UUID, p.Name, p.FirstName, p.LastName, p.GenerationName, p.Alias, 
			p.EnglishName, p.Gender, p.IsAlive, p.BirthDate, p.BirthLunar, 
			p.BirthPlace, p.BirthPlaceLongitude, p.BirthPlaceLatitude, 
			p.DeathDate, p.DeathLunar, p.DeathPlace, p.DeathPlaceLongitude, 
			p.DeathPlaceLatitude, p.DeathCause, p.Generation, p.Occupation, 
			p.Education, p.Biography, p.AvatarURL, p.CoverPhotoURL, p.IsPublic, 
			p.VerificationStatus, p.CreatedBy, p.CreatedAt, p.UpdatedAt,
		).
		Build()

	_, err := SafeExec(r.db, query, args...)
	return err
}

// Update 更新人员信息
func (r *PersonRepository) Update(p *person.Person) error {
	builder := NewUpdateBuilder("person")
	query, args := builder.
		Set("uuid", p.UUID).
		Set("name", p.Name).
		Set("first_name", p.FirstName).
		Set("last_name", p.LastName).
		Set("generation_name", p.GenerationName).
		Set("alias", p.Alias).
		Set("english_name", p.EnglishName).
		Set("gender", p.Gender).
		Set("is_alive", p.IsAlive).
		Set("birth_date", p.BirthDate).
		Set("birth_lunar", p.BirthLunar).
		Set("birth_place", p.BirthPlace).
		Set("birth_place_longitude", p.BirthPlaceLongitude).
		Set("birth_place_latitude", p.BirthPlaceLatitude).
		Set("death_date", p.DeathDate).
		Set("death_lunar", p.DeathLunar).
		Set("death_place", p.DeathPlace).
		Set("death_place_longitude", p.DeathPlaceLongitude).
		Set("death_place_latitude", p.DeathPlaceLatitude).
		Set("death_cause", p.DeathCause).
		Set("generation", p.Generation).
		Set("occupation", p.Occupation).
		Set("education", p.Education).
		Set("biography", p.Biography).
		Set("avatar_url", p.AvatarURL).
		Set("cover_photo_url", p.CoverPhotoURL).
		Set("is_public", p.IsPublic).
		Set("verification_status", p.VerificationStatus).
		Set("created_by", p.CreatedBy).
		Set("updated_at", time.Now()).
		Where("id = ?", p.ID).
		Build()

	_, err := SafeExec(r.db, query, args...)
	return err
}

// Delete 删除人员
func (r *PersonRepository) Delete(id int) error {
	builder := NewDeleteBuilder("person")
	query, args := builder.Where("id = ?", id).Build()
	
	_, err := SafeExec(r.db, query, args...)
	return err
}

// GetByID 根据ID获取人员
func (r *PersonRepository) GetByID(id int) (*person.Person, error) {
	builder := NewSQLBuilder()
	query, args := builder.
		Select(
			"id", "uuid", "name", "first_name", "last_name", "generation_name", 
			"alias", "english_name", "gender", "is_alive", "birth_date", 
			"birth_lunar", "birth_place", "birth_place_longitude", 
			"birth_place_latitude", "death_date", "death_lunar", "death_place", 
			"death_place_longitude", "death_place_latitude", "death_cause", 
			"generation", "occupation", "education", "biography", "avatar_url", 
			"cover_photo_url", "is_public", "verification_status", "created_by", 
			"created_at", "updated_at",
		).
		From("person").
		Where("id = ?", id).
		Build()

	row := SafeQueryRow(r.db, query, args...)

	p := &person.Person{}
	err := row.Scan(
		&p.ID, &p.UUID, &p.Name, &p.FirstName, &p.LastName, &p.GenerationName, 
		&p.Alias, &p.EnglishName, &p.Gender, &p.IsAlive, &p.BirthDate, 
		&p.BirthLunar, &p.BirthPlace, &p.BirthPlaceLongitude, &p.BirthPlaceLatitude, 
		&p.DeathDate, &p.DeathLunar, &p.DeathPlace, &p.DeathPlaceLongitude, 
		&p.DeathPlaceLatitude, &p.DeathCause, &p.Generation, &p.Occupation, 
		&p.Education, &p.Biography, &p.AvatarURL, &p.CoverPhotoURL, &p.IsPublic, 
		&p.VerificationStatus, &p.CreatedBy, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// GetByUUID 根据UUID获取人员
func (r *PersonRepository) GetByUUID(uuid string) (*person.Person, error) {
	builder := NewSQLBuilder()
	query, args := builder.
		Select(
			"id", "uuid", "name", "first_name", "last_name", "generation_name", 
			"alias", "english_name", "gender", "is_alive", "birth_date", 
			"birth_lunar", "birth_place", "birth_place_longitude", 
			"birth_place_latitude", "death_date", "death_lunar", "death_place", 
			"death_place_longitude", "death_place_latitude", "death_cause", 
			"generation", "occupation", "education", "biography", "avatar_url", 
			"cover_photo_url", "is_public", "verification_status", "created_by", 
			"created_at", "updated_at",
		).
		From("person").
		Where("uuid = ?", uuid).
		Build()

	row := SafeQueryRow(r.db, query, args...)

	p := &person.Person{}
	err := row.Scan(
		&p.ID, &p.UUID, &p.Name, &p.FirstName, &p.LastName, &p.GenerationName, 
		&p.Alias, &p.EnglishName, &p.Gender, &p.IsAlive, &p.BirthDate, 
		&p.BirthLunar, &p.BirthPlace, &p.BirthPlaceLongitude, &p.BirthPlaceLatitude, 
		&p.DeathDate, &p.DeathLunar, &p.DeathPlace, &p.DeathPlaceLongitude, 
		&p.DeathPlaceLatitude, &p.DeathCause, &p.Generation, &p.Occupation, 
		&p.Education, &p.Biography, &p.AvatarURL, &p.CoverPhotoURL, &p.IsPublic, 
		&p.VerificationStatus, &p.CreatedBy, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// List 列出人员
func (r *PersonRepository) List(page, pageSize int) ([]*person.Person, int, error) {
	offset := (page - 1) * pageSize

	// 获取总数
	countBuilder := NewSQLBuilder()
	countQuery, countArgs := countBuilder.Select("COUNT(*)").From("person").Build()
	
	var total int
	err := SafeQueryRow(r.db, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	builder := NewSQLBuilder()
	query, args := builder.
		Select(
			"id", "uuid", "name", "first_name", "last_name", "generation_name", 
			"alias", "english_name", "gender", "is_alive", "birth_date", 
			"birth_lunar", "birth_place", "birth_place_longitude", 
			"birth_place_latitude", "death_date", "death_lunar", "death_place", 
			"death_place_longitude", "death_place_latitude", "death_cause", 
			"generation", "occupation", "education", "biography", "avatar_url", 
			"cover_photo_url", "is_public", "verification_status", "created_by", 
			"created_at", "updated_at",
		).
		From("person").
		Limit(pageSize).
		Offset(offset).
		Build()

	rows, err := SafeQuery(r.db, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	persons := []*person.Person{}
	for rows.Next() {
		p := &person.Person{}
		err := rows.Scan(
			&p.ID, &p.UUID, &p.Name, &p.FirstName, &p.LastName, &p.GenerationName, 
			&p.Alias, &p.EnglishName, &p.Gender, &p.IsAlive, &p.BirthDate, 
			&p.BirthLunar, &p.BirthPlace, &p.BirthPlaceLongitude, &p.BirthPlaceLatitude, 
			&p.DeathDate, &p.DeathLunar, &p.DeathPlace, &p.DeathPlaceLongitude, 
			&p.DeathPlaceLatitude, &p.DeathCause, &p.Generation, &p.Occupation, 
			&p.Education, &p.Biography, &p.AvatarURL, &p.CoverPhotoURL, &p.IsPublic, 
			&p.VerificationStatus, &p.CreatedBy, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		persons = append(persons, p)
	}

	return persons, total, nil
}

// Search 搜索人员
func (r *PersonRepository) Search(name string, page, pageSize int) ([]*person.Person, int, error) {
	offset := (page - 1) * pageSize

	// 获取总数
	countBuilder := NewSQLBuilder()
	countQuery, countArgs := countBuilder.
		Select("COUNT(*)").
		From("person").
		Like("name", name).
		Build()
	
	var total int
	err := SafeQueryRow(r.db, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	builder := NewSQLBuilder()
	query, args := builder.
		Select(
			"id", "uuid", "name", "first_name", "last_name", "generation_name", 
			"alias", "english_name", "gender", "is_alive", "birth_date", 
			"birth_lunar", "birth_place", "birth_place_longitude", 
			"birth_place_latitude", "death_date", "death_lunar", "death_place", 
			"death_place_longitude", "death_place_latitude", "death_cause", 
			"generation", "occupation", "education", "biography", "avatar_url", 
			"cover_photo_url", "is_public", "verification_status", "created_by", 
			"created_at", "updated_at",
		).
		From("person").
		Like("name", name).
		Limit(pageSize).
		Offset(offset).
		Build()

	rows, err := SafeQuery(r.db, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	persons := []*person.Person{}
	for rows.Next() {
		p := &person.Person{}
		err := rows.Scan(
			&p.ID, &p.UUID, &p.Name, &p.FirstName, &p.LastName, &p.GenerationName, 
			&p.Alias, &p.EnglishName, &p.Gender, &p.IsAlive, &p.BirthDate, 
			&p.BirthLunar, &p.BirthPlace, &p.BirthPlaceLongitude, &p.BirthPlaceLatitude, 
			&p.DeathDate, &p.DeathLunar, &p.DeathPlace, &p.DeathPlaceLongitude, 
			&p.DeathPlaceLatitude, &p.DeathCause, &p.Generation, &p.Occupation, 
			&p.Education, &p.Biography, &p.AvatarURL, &p.CoverPhotoURL, &p.IsPublic, 
			&p.VerificationStatus, &p.CreatedBy, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		persons = append(persons, p)
	}

	return persons, total, nil
}
