package db

import (
	"GenealogyManagementSystem/internal/domain/person"
	"GenealogyManagementSystem/internal/infrastructure/persistence/sqlite"
	"database/sql"
	"os"
	"testing"
	"time"

	_ "modernc.org/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestPersonRepository_BasicOperations(t *testing.T) {
	// 创建临时数据库文件
	testDBPath := "test_genealogy.db"
	defer os.Remove(testDBPath)

	// 初始化数据库连接
	db, err := sql.Open("sqlite", testDBPath)
	assert.NoError(t, err)
	defer db.Close()

	// 创建表结构
	err = createTables(db)
	assert.NoError(t, err)

	// 创建仓储实例
	repo := sqlite.NewPersonRepository(db)

	// 测试创建人员
	p := &person.Person{
		UUID:        "test-uuid-123",
		Name:        "张三",
		FirstName:   "张",
		LastName:    "三",
		Gender:      "男",
		IsAlive:     true,
		BirthDate:   time.Now(),
		BirthPlace:  "北京",
		Generation:  1,
		IsPublic:    true,
		CreatedBy:   1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = repo.Create(p)
	assert.NoError(t, err)

	// 测试列表查询
	result, total, err := repo.List(1, 10)
	assert.NoError(t, err)
	assert.Equal(t, 1, total)
	assert.Len(t, result, 1)
	assert.Equal(t, "张三", result[0].Name)
}

func TestPersonRepository_UpdateAndDelete(t *testing.T) {
	// 创建临时数据库文件
	testDBPath := "test_genealogy_operations.db"
	defer os.Remove(testDBPath)

	// 初始化数据库连接
	db, err := sql.Open("sqlite", testDBPath)
	assert.NoError(t, err)
	defer db.Close()

	// 创建表结构
	err = createTables(db)
	assert.NoError(t, err)

	// 创建仓储实例
	repo := sqlite.NewPersonRepository(db)

	// 创建测试数据
	p := &person.Person{
		UUID:      "operations-test-uuid",
		Name:      "原始姓名",
		Gender:    "男",
		IsAlive:   true,
		BirthDate: time.Now(),
		CreatedBy: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = repo.Create(p)
	assert.NoError(t, err)

	// 获取插入的记录ID
	result, total, err := repo.List(1, 10)
	assert.NoError(t, err)
	assert.Equal(t, 1, total)
	
	insertedPerson := result[0]
	
	// 测试更新
	insertedPerson.Name = "更新后的姓名"
	err = repo.Update(insertedPerson)
	assert.NoError(t, err)

	// 验证更新
	updatedResult, _, err := repo.List(1, 10)
	assert.NoError(t, err)
	assert.Equal(t, "更新后的姓名", updatedResult[0].Name)

	// 测试删除
	err = repo.Delete(insertedPerson.ID)
	assert.NoError(t, err)

	// 验证删除
	finalResult, finalTotal, err := repo.List(1, 10)
	assert.NoError(t, err)
	assert.Equal(t, 0, finalTotal)
	assert.Len(t, finalResult, 0)
}

func TestPersonRepository_GetByUUID(t *testing.T) {
	// 创建临时数据库文件
	testDBPath := "test_genealogy_uuid.db"
	defer os.Remove(testDBPath)

	// 初始化数据库连接
	db, err := sql.Open("sqlite", testDBPath)
	assert.NoError(t, err)
	defer db.Close()

	// 创建表结构
	err = createTables(db)
	assert.NoError(t, err)

	// 创建仓储实例
	repo := sqlite.NewPersonRepository(db)

	// 创建测试数据
	p := &person.Person{
		UUID:      "uuid-test-123",
		Name:      "UUID测试用户",
		Gender:    "女",
		IsAlive:   true,
		BirthDate: time.Now(),
		CreatedBy: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = repo.Create(p)
	assert.NoError(t, err)

	// 测试通过UUID获取
	retrieved, err := repo.GetByUUID("uuid-test-123")
	assert.NoError(t, err)
	assert.Equal(t, "UUID测试用户", retrieved.Name)
	assert.Equal(t, "uuid-test-123", retrieved.UUID)
}

func TestPersonRepository_NotFound(t *testing.T) {
	// 创建临时数据库文件
	testDBPath := "test_genealogy_notfound.db"
	defer os.Remove(testDBPath)

	// 初始化数据库连接
	db, err := sql.Open("sqlite", testDBPath)
	assert.NoError(t, err)
	defer db.Close()

	// 创建表结构
	err = createTables(db)
	assert.NoError(t, err)

	// 创建仓储实例
	repo := sqlite.NewPersonRepository(db)

	// 测试获取不存在的记录
	retrieved, err := repo.GetByID(99999)
	assert.Error(t, err)
	assert.Nil(t, retrieved)

	retrieved, err = repo.GetByUUID("non-existent-uuid")
	assert.Error(t, err)
	assert.Nil(t, retrieved)
}

func createTables(db *sql.DB) error {
	createTableSQL := `
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

	_, err := db.Exec(createTableSQL)
	return err
}