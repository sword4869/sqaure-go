package store

import (
	"database/sql"
	"test/config"
)

type PostLike struct {
	Id        int   `db:"id" json:"id"`
	UserId    int   `db:"user_id" json:"user_id"`
	PostId    int   `db:"post_id" json:"post_id"`
	IsLike    int   `db:"is_like" json:"is_like"`
	CreatedAt int64 `db:"created_at" json:"created_at"`
}

func NewPostLike() *PostLike {
	return &PostLike{}
}

func (p *PostLike) Create() error {
	query := `INSERT INTO post_likes (user_id, post_id, is_like, created_at) VALUES (?, ?, ?, ?)`
	_, err := config.Db.Exec(query, p.UserId, p.PostId, p.IsLike, p.CreatedAt)
	if err != nil {
		return err
	}
	p.Id = int(p.Id)
	return nil
}

func (p *PostLike) GetByUserIdAndPostId(userId int, postId int) (*PostLike, error) {
	postLike := &PostLike{}
	query := `SELECT * FROM post_likes WHERE user_id = ? AND post_id = ?`
	err := config.Db.Unsafe().Get(postLike, query, userId, postId)
	if err != nil && err != sql.ErrNoRows {
		return postLike, err
	}
	return postLike, nil
}

func (p *PostLike) UpdateLiked(userId int, postId int, isLike int) error {
	query := `REPLACE INTO post_likes (user_id, post_id, is_like) VALUES (?, ?, ?)`
	_, err := config.Db.Exec(query, userId, postId, isLike)
	if err != nil {
		return err
	}
	return nil
}
