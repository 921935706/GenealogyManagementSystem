package sqlite

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB 初始化数据库连接并创建表结构
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("sqlite", dataSourceName)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	// 创建表结构
	err = createTables()
	if err != nil {
		return err
	}

	log.Println("Database connected and tables created successfully")
	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// createTables 创建所有必要的表结构
func createTables() error {
	// 人员表
	personTableSQL := `
	CREATE TABLE IF NOT EXISTS person (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL,
		first_name TEXT,
		last_name TEXT,
		generation_name TEXT,
		alias TEXT,
		english_name TEXT,
		gender TEXT,
		is_alive BOOLEAN DEFAULT true,
		birth_date DATETIME,
		birth_lunar TEXT,
		birth_place TEXT,
		birth_place_longitude REAL,
		birth_place_latitude REAL,
		death_date DATETIME,
		death_lunar TEXT,
		death_place TEXT,
		death_place_longitude REAL,
		death_place_latitude REAL,
		death_cause TEXT,
		generation INTEGER,
		occupation TEXT,
		education TEXT,
		biography TEXT,
		avatar_url TEXT,
		cover_photo_url TEXT,
		is_public BOOLEAN DEFAULT true,
		verification_status TEXT DEFAULT 'pending',
		created_by INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	// 关系表
	relationshipTableSQL := `
	CREATE TABLE IF NOT EXISTS relationship (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		parent_id INTEGER NOT NULL,
		child_id INTEGER NOT NULL,
		relationship_type TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (parent_id) REFERENCES person(id),
		FOREIGN KEY (child_id) REFERENCES person(id),
		UNIQUE(parent_id, child_id)
	)
	`

	// 家庭关系表
	familyRelationTableSQL := `
	CREATE TABLE IF NOT EXISTS family_relation (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		person_id INTEGER NOT NULL,
		spouse_id INTEGER,
		family_id INTEGER,
		relation_type TEXT NOT NULL,
		marriage_date DATETIME,
		divorce_date DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (person_id) REFERENCES person(id),
		FOREIGN KEY (spouse_id) REFERENCES person(id)
	)
	`

	// 事件表
	eventTableSQL := `
	CREATE TABLE IF NOT EXISTS event (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		person_id INTEGER NOT NULL,
		event_type TEXT NOT NULL,
		event_date DATETIME NOT NULL,
		event_place TEXT,
		event_description TEXT,
		is_important BOOLEAN DEFAULT false,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (person_id) REFERENCES person(id)
	)
	`

	// 多媒体表
	mediaTableSQL := `
	CREATE TABLE IF NOT EXISTS media (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		person_id INTEGER,
		event_id INTEGER,
		media_type TEXT NOT NULL,
		file_url TEXT NOT NULL,
		file_name TEXT,
		description TEXT,
		is_public BOOLEAN DEFAULT true,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (person_id) REFERENCES person(id),
		FOREIGN KEY (event_id) REFERENCES event(id)
	)
	`

	tables := []string{
		personTableSQL,
		relationshipTableSQL,
		familyRelationTableSQL,
		eventTableSQL,
		mediaTableSQL,
	}

	for _, tableSQL := range tables {
		_, err := DB.Exec(tableSQL)
		if err != nil {
			return err
		}
	}

	return nil
}
