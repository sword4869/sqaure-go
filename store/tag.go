package store

import (
	"database/sql"
	"test/config"
)

// Tag 标签模型
type Tag struct {
	Id        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Color     string `db:"color" json:"color"`
	IsActive  int    `db:"is_active" json:"is_active"`
	SortOrder int    `db:"sort_order" json:"sort_order"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}

func NewTag() *Tag {
	return &Tag{}
}

// GetByID 根据ID获取标签
func (t *Tag) GetByID(id int) (*Tag, error) {
	tag := &Tag{}
	query := `SELECT * FROM tags WHERE id = ?`

	err := config.Db.Unsafe().Get(tag, query, id)
	if err != nil && err != sql.ErrNoRows {
		return tag, err
	}
	return tag, nil
}

// ListAllTags 获取所有标签
func (t *Tag) ListAllTags() ([]*Tag, error) {
	var tags []*Tag
	query := `SELECT * FROM tags WHERE is_active = 1 ORDER BY sort_order ASC, id ASC`

	err := config.Db.Unsafe().Select(&tags, query)
	if err != nil && err != sql.ErrNoRows {
		return tags, err
	}

	return tags, nil
}
