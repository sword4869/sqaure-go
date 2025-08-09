package store

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"test/config"
	"time"
)

// Post 发帖模型
type Post struct {
	Id        int      `db:"id" json:"id" example:"1"`
	UserId    int      `db:"user_id" json:"user_id" binding:"required" example:"1"`
	PostTags  PostTags `db:"post_tags" json:"post_tags" binding:"required" example:"1,2,3"`
	Content   string   `db:"content" json:"content" binding:"required" example:"1"`
	Images    Images   `db:"images" json:"images" example:"1,2"`
	Latitude  float64  `db:"latitude" json:"latitude" example:"120.123456"`
	Longitude float64  `db:"longitude" json:"longitude" example:"30.123456"`
	IsActive  int      `db:"is_active" json:"is_active" example:"1"` // 是否展示 0 不展示 1 展示
	LikeCount int      `db:"like_count" json:"like_count" example:"0"`
	CreatedAt int64    `db:"created_at" json:"created_at"`
	UpdatedAt int64    `db:"updated_at" json:"updated_at"`
}

func NewPost() *Post {
	return &Post{
		Images: make(Images, 0),
	}
}

type Images []int64

func (i *Images) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), i)
}

func (i Images) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type PostTags []int

func (i *PostTags) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), i)
}

func (i PostTags) Value() (driver.Value, error) {
	return json.Marshal(i)
}

// Create 创建发帖
func (p *Post) Create() error {
	// 获取当前时间戳
	now := time.Now().Unix()
	p.CreatedAt = now
	p.UpdatedAt = now
	// 默认是展示
	p.IsActive = 1

	query := `INSERT INTO posts (user_id, post_tags, content, images, latitude, longitude, is_active, like_count, created_at, updated_at)  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := config.Db.Exec(query, p.UserId, p.PostTags, p.Content, p.Images, p.Latitude, p.Longitude, p.IsActive, p.LikeCount, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.Id = int(id)
	return nil
}

func (p *Post) GetById(id int) (*Post, error) {
	post := &Post{}
	query := `SELECT * FROM posts WHERE id = ?`

	err := config.Db.Unsafe().Get(post, query, id)
	if err != nil && err != sql.ErrNoRows {
		return post, err
	}
	return post, nil
}

// ListPosts 获取发帖列表 cursor 游标
func (p *Post) ListPosts(cursor int64) ([]*Post, error) {
	var posts []*Post
	query := `SELECT * FROM posts WHERE is_active = 1 and created_at < ? ORDER BY created_at DESC LIMIT 10`
	err := config.Db.Unsafe().Select(&posts, query, cursor)
	if err != nil && err != sql.ErrNoRows {
		return posts, err
	}

	return posts, nil
}

// ListPostsByKeyword 通过关键字获取发帖列表
func (p *Post) ListPostsByKeyword(keyword string, cursor int64) ([]*Post, error) {
	var posts []*Post
	query := `SELECT * FROM posts WHERE content LIKE ? AND is_active = 1 and created_at < ? ORDER BY created_at DESC LIMIT 10`
	err := config.Db.Unsafe().Select(&posts, query, "%"+keyword+"%", cursor)
	if err != nil && err != sql.ErrNoRows {
		return posts, err
	}
	return posts, nil
}

// Update 更新发帖
func (p *Post) Update() error {
	p.UpdatedAt = time.Now().Unix()
	query := `UPDATE posts SET content = ?, images = ?, latitude = ?, longitude = ?, like_count = ?, updated_at = ? WHERE id = ?`

	_, err := config.Db.Exec(query, p.Content, p.Images,
		p.Latitude, p.Longitude, p.LikeCount, p.UpdatedAt, p.Id)
	return err
}

// Like 点赞
func (p *Post) Like(id int) error {
	query := `UPDATE posts SET like_count = like_count + 1 WHERE id = ?`
	_, err := config.Db.Exec(query, id)
	return err
}

// UnLike 取消点赞
func (p *Post) UnLike(id int) error {
	query := `UPDATE posts SET like_count = like_count - 1 WHERE id = ?`
	_, err := config.Db.Exec(query, id)
	return err
}

// Delete 删除发帖
func (p *Post) Delete(id int) error {
	query := `DELETE FROM posts WHERE id = ?`
	_, err := config.Db.Exec(query, id)
	return err
}
