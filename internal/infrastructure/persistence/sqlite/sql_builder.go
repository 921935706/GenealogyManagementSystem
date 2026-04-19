package sqlite

import (
	"database/sql"
	"fmt"
	"strings"
)

// SQLBuilder SQL构建器，提供安全的SQL语句构建功能
type SQLBuilder struct {
	query      strings.Builder
	args       []interface{}
	hasWhere   bool
	hasOrderBy bool
}

// NewSQLBuilder 创建新的SQL构建器
func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{
		args: make([]interface{}, 0),
	}
}

// Select 添加SELECT子句
func (b *SQLBuilder) Select(columns ...string) *SQLBuilder {
	b.query.WriteString("SELECT ")
	b.query.WriteString(strings.Join(columns, ", "))
	return b
}

// From 添加FROM子句
func (b *SQLBuilder) From(table string) *SQLBuilder {
	b.query.WriteString(" FROM ")
	b.query.WriteString(table)
	return b
}

// Where 添加WHERE条件
func (b *SQLBuilder) Where(condition string, args ...interface{}) *SQLBuilder {
	if !b.hasWhere {
		b.query.WriteString(" WHERE ")
		b.hasWhere = true
	} else {
		b.query.WriteString(" AND ")
	}
	b.query.WriteString(condition)
	b.args = append(b.args, args...)
	return b
}

// OrWhere 添加OR WHERE条件
func (b *SQLBuilder) OrWhere(condition string, args ...interface{}) *SQLBuilder {
	if !b.hasWhere {
		b.query.WriteString(" WHERE ")
		b.hasWhere = true
	} else {
		b.query.WriteString(" OR ")
	}
	b.query.WriteString(condition)
	b.args = append(b.args, args...)
	return b
}

// Like 添加LIKE条件，安全处理通配符
func (b *SQLBuilder) Like(column, value string) *SQLBuilder {
	return b.Where(fmt.Sprintf("%s LIKE ?", column), "%"+escapeLike(value)+"%")
}

// OrderBy 添加ORDER BY子句
func (b *SQLBuilder) OrderBy(column string, direction string) *SQLBuilder {
	if !b.hasOrderBy {
		b.query.WriteString(" ORDER BY ")
		b.hasOrderBy = true
	} else {
		b.query.WriteString(", ")
	}
	b.query.WriteString(column)
	b.query.WriteString(" ")
	b.query.WriteString(direction)
	return b
}

// Limit 添加LIMIT子句
func (b *SQLBuilder) Limit(limit int) *SQLBuilder {
	b.query.WriteString(" LIMIT ?")
	b.args = append(b.args, limit)
	return b
}

// Offset 添加OFFSET子句
func (b *SQLBuilder) Offset(offset int) *SQLBuilder {
	b.query.WriteString(" OFFSET ?")
	b.args = append(b.args, offset)
	return b
}

// Build 构建SQL语句和参数
func (b *SQLBuilder) Build() (string, []interface{}) {
	return b.query.String(), b.args
}

// InsertBuilder INSERT语句构建器
type InsertBuilder struct {
	table   string
	columns []string
	values  []interface{}
}

// NewInsertBuilder 创建INSERT构建器
func NewInsertBuilder(table string) *InsertBuilder {
	return &InsertBuilder{
		table: table,
	}
}

// Columns 设置插入的列
func (b *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	b.columns = columns
	return b
}

// Values 设置插入的值
func (b *InsertBuilder) Values(values ...interface{}) *InsertBuilder {
	b.values = values
	return b
}

// Build 构建INSERT语句
func (b *InsertBuilder) Build() (string, []interface{}) {
	placeholders := make([]string, len(b.columns))
	for i := range b.columns {
		placeholders[i] = "?"
	}
	
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		b.table,
		strings.Join(b.columns, ", "),
		strings.Join(placeholders, ", "),
	)
	
	return query, b.values
}

// UpdateBuilder UPDATE语句构建器
type UpdateBuilder struct {
	table     string
	setValues map[string]interface{}
	where     string
	whereArgs []interface{}
}

// NewUpdateBuilder 创建UPDATE构建器
func NewUpdateBuilder(table string) *UpdateBuilder {
	return &UpdateBuilder{
		table:     table,
		setValues: make(map[string]interface{}),
	}
}

// Set 设置更新的列和值
func (b *UpdateBuilder) Set(column string, value interface{}) *UpdateBuilder {
	b.setValues[column] = value
	return b
}

// Where 设置WHERE条件
func (b *UpdateBuilder) Where(condition string, args ...interface{}) *UpdateBuilder {
	b.where = condition
	b.whereArgs = args
	return b
}

// Build 构建UPDATE语句
func (b *UpdateBuilder) Build() (string, []interface{}) {
	setClauses := make([]string, 0, len(b.setValues))
	args := make([]interface{}, 0, len(b.setValues))
	
	for column, value := range b.setValues {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", column))
		args = append(args, value)
	}
	
	query := fmt.Sprintf("UPDATE %s SET %s", b.table, strings.Join(setClauses, ", "))
	
	if b.where != "" {
		query += " WHERE " + b.where
		args = append(args, b.whereArgs...)
	}
	
	return query, args
}

// DeleteBuilder DELETE语句构建器
type DeleteBuilder struct {
	table     string
	where     string
	whereArgs []interface{}
}

// NewDeleteBuilder 创建DELETE构建器
func NewDeleteBuilder(table string) *DeleteBuilder {
	return &DeleteBuilder{
		table: table,
	}
}

// Where 设置WHERE条件
func (b *DeleteBuilder) Where(condition string, args ...interface{}) *DeleteBuilder {
	b.where = condition
	b.whereArgs = args
	return b
}

// Build 构建DELETE语句
func (b *DeleteBuilder) Build() (string, []interface{}) {
	query := fmt.Sprintf("DELETE FROM %s", b.table)
	args := make([]interface{}, 0)
	
	if b.where != "" {
		query += " WHERE " + b.where
		args = append(args, b.whereArgs...)
	}
	
	return query, args
}

// escapeLike 转义LIKE查询中的特殊字符
func escapeLike(value string) string {
	// 转义SQL LIKE中的特殊字符：%, _, \
	replacer := strings.NewReplacer(
		"\\", "\\\\",
		"%", "\\%",
		"_", "\\_",
	)
	return replacer.Replace(value)
}

// SafeQuery 安全的查询执行函数
func SafeQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	// 可以在这里添加额外的安全检查
	return db.Query(query, args...)
}

// SafeQueryRow 安全的单行查询执行函数
func SafeQueryRow(db *sql.DB, query string, args ...interface{}) *sql.Row {
	// 可以在这里添加额外的安全检查
	return db.QueryRow(query, args...)
}

// SafeExec 安全的执行函数
func SafeExec(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	// 可以在这里添加额外的安全检查
	return db.Exec(query, args...)
}