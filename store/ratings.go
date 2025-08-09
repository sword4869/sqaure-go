package store

import (
	"database/sql"
	"test/config"
	"time"
)

type Ratings struct {
	Id        int   `db:"id" json:"id"`
	UserId    int   `db:"user_id" json:"user_id" binding:"required"` // 用户ID
	PostId    int   `db:"post_id" json:"post_id" binding:"required"` // 帖子ID
	Rating    int   `db:"rating" json:"rating" binding:"required"`   // 评分
	CreatedAt int64 `db:"created_at" json:"created_at"`
	UpdatedAt int64 `db:"updated_at" json:"updated_at"`
}

func NewRatings() *Ratings {
	return &Ratings{}
}

func (s *Ratings) Create() error {
	now := time.Now().Unix()
	query := `INSERT INTO ratings (user_id, post_id, rating, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_, err := config.Db.Exec(query, s.UserId, s.PostId, s.Rating, now, now)
	if err != nil {
		return err
	}
	s.Id = int(s.Id)
	return nil
}

func (s *Ratings) GetByUserIdAndPostId(userId int, postId int) (*Ratings, error) {
	ratings := &Ratings{}
	query := `SELECT * FROM ratings WHERE user_id = ? AND post_id = ?`
	err := config.Db.Unsafe().Get(ratings, query, userId, postId)
	if err != nil && err != sql.ErrNoRows {
		return ratings, err
	}
	return ratings, nil
}
