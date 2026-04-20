# 族谱管理系统 - 数据模型与 API 设计文档

## 一、技术栈说明

- 数据库：SQLite
---

## 二、数据库表结构
存储每个家族成员的基本信息。

```sql
CREATE TABLE person (
    -- 基础标识
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid VARCHAR(36) NOT NULL UNIQUE,              -- 全局唯一标识，用于导入导出
    
    -- 姓名相关
    name VARCHAR(50) NOT NULL,                     -- 常用名
    first_name VARCHAR(50),                        -- 名
    last_name VARCHAR(50),                         -- 姓
    generation_name VARCHAR(50),                   -- 字辈名（如"庆"字辈）
    alias VARCHAR(100),                            -- 别名/字/号
    english_name VARCHAR(100),                     -- 英文名
    
    -- 性别与状态
    gender TEXT DEFAULT 'unknown',                 -- SQLite不支持ENUM，使用TEXT代替
    is_alive BOOLEAN DEFAULT TRUE,                 -- 是否在世
    
    -- 出生信息
    birth_date DATE,                               -- 出生日期（公历）
    birth_lunar VARCHAR(20),                       -- 农历出生日期
    birth_place VARCHAR(200),                      -- 出生地点
    birth_place_longitude DECIMAL(10,7),           -- 经度
    birth_place_latitude DECIMAL(10,7),            -- 纬度
    
    -- 死亡信息
    death_date DATE,
    death_lunar VARCHAR(20),
    death_place VARCHAR(200),
    death_place_longitude DECIMAL(10,7),
    death_place_latitude DECIMAL(10,7),
    death_cause TEXT,                              -- 死因
    
    -- 其他信息
    generation INT,                                -- 世代编号（从始祖为1开始）
    occupation VARCHAR(100),                       -- 职业
    education VARCHAR(100),                        -- 学历
    biography TEXT,                                -- 生平简介（富文本）
    
    -- 多媒体
    avatar_url VARCHAR(500),                       -- 头像URL
    cover_photo_url VARCHAR(500),                  -- 封面照片
    
    -- 状态字段
    is_public BOOLEAN DEFAULT FALSE,               -- 是否公开
    verification_status TEXT DEFAULT 'pending',    -- SQLite不支持ENUM，使用TEXT代替
    created_by INTEGER,                            -- 创建人ID
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    INDEX idx_name (name),
    INDEX idx_generation (generation),
    INDEX idx_birth_date (birth_date),
    INDEX idx_is_alive (is_alive)
);
```

###  2.2 家庭关系表（family_relation） 
记录夫妻/伴侣关系，支持再婚、一夫多妻等复杂情况。

```sql
CREATE TABLE family_relation (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    family_id VARCHAR(36) NOT NULL,                -- 家庭唯一标识
    person_id INTEGER NOT NULL,                    -- 人员ID
    spouse_id INTEGER NOT NULL,                    -- 配偶ID
    marriage_date DATE,                            -- 结婚日期
    marriage_place VARCHAR(200),                   -- 结婚地点
    divorce_date DATE,                             -- 离婚日期
    relation_type TEXT DEFAULT 'married',          -- SQLite不支持ENUM，使用TEXT代替
    is_primary BOOLEAN DEFAULT FALSE,              -- 是否为主要配偶
    notes TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (person_id) REFERENCES person(id) ON DELETE CASCADE,
    FOREIGN KEY (spouse_id) REFERENCES person(id) ON DELETE CASCADE,
    UNIQUE (family_id, person_id, spouse_id),      -- SQLite使用UNIQUE约束
    INDEX idx_person (person_id),
    INDEX idx_spouse (spouse_id)
);
```
###  2.3 亲子关系表（parent_child）
记录父母与子女的关系，支持继父母、养父母。

```sql

CREATE TABLE parent_child (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    parent_id INTEGER NOT NULL,                    -- 父母ID
    child_id INTEGER NOT NULL,                     -- 子女ID
    relation_type TEXT DEFAULT 'biological',       -- SQLite不支持ENUM，使用TEXT代替
    is_primary BOOLEAN DEFAULT TRUE,               -- 是否为主要监护人
    notes TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (parent_id) REFERENCES person(id) ON DELETE CASCADE,
    FOREIGN KEY (child_id) REFERENCES person(id) ON DELETE CASCADE,
    UNIQUE (parent_id, child_id),                  -- SQLite使用UNIQUE约束
    INDEX idx_parent (parent_id),
    INDEX idx_child (child_id)
);

```
### 2.4 世系表（pedigree）
存储家族树结构，用于快速查询上下几代人。

```sql
CREATE TABLE pedigree (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    person_id INTEGER NOT NULL,
    father_id INTEGER,                             -- 父亲ID（冗余字段）
    mother_id INTEGER,                             -- 母亲ID
    spouse_id INTEGER,                             -- 当前配偶ID（冗余）
    generation INT,                                -- 世代数（从始祖0开始）
    lineage_path VARCHAR(1000),                    -- 世系路径（如"1-3-5-8"）
    depth INT DEFAULT 0,                           -- 深度（从根节点开始的层数）
    
    FOREIGN KEY (person_id) REFERENCES person(id) ON DELETE CASCADE,
    INDEX idx_generation (generation),
    INDEX idx_path (lineage_path)
);

```

### 2.5 事件表（event）
记录家族成员的各类事件（毕业、升职、迁徙等）。

```sql
CREATE TABLE event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    person_id INTEGER NOT NULL,
    event_type VARCHAR(50) NOT NULL,               -- 事件类型
    event_name VARCHAR(200),                       -- 事件名称
    event_date DATE,
    event_lunar VARCHAR(20),
    event_place VARCHAR(200),
    event_place_longitude DECIMAL(10,7),
    event_place_latitude DECIMAL(10,7),
    description TEXT,
    is_important BOOLEAN DEFAULT FALSE,            -- 是否重要事件
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (person_id) REFERENCES person(id) ON DELETE CASCADE,
    INDEX idx_person (person_id),
    INDEX idx_event_date (event_date),
    INDEX idx_event_type (event_type)
);
```

### 2.6 多媒体表（media）
存储照片、文档、音频、视频等。

```sql
CREATE TABLE media (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    person_id INTEGER,                             -- 关联的人员
    event_id INTEGER,                              -- 关联的事件
    media_type TEXT DEFAULT 'photo',               -- SQLite不支持ENUM，使用TEXT代替
    file_name VARCHAR(255) NOT NULL,
    file_url VARCHAR(500) NOT NULL,
    file_size INTEGER,
    mime_type VARCHAR(100),
    thumbnail_url VARCHAR(500),
    title VARCHAR(200),
    description TEXT,
    uploader_id INTEGER,
    is_public BOOLEAN DEFAULT FALSE,
    uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (person_id) REFERENCES person(id) ON DELETE SET NULL,
    INDEX idx_person (person_id),
    INDEX idx_media_type (media_type)
);
```