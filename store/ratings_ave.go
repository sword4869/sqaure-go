package store

import (
	"database/sql"
	"test/config"
)

type RatingsAve struct {
	PostId      int     `db:"post_id" json:"post_id"`           // 发帖ID
	RatingAve   float64 `db:"rating_ave" json:"rating_ave"`     // 评分平均值
	RatingCount int     `db:"rating_count" json:"rating_count"` // 评分总人数
}

func NewRatingsAve() *RatingsAve {
	return &RatingsAve{}
}

func NewRatingsAveByPostId(postId int) *RatingsAve {
	return &RatingsAve{
		PostId: postId,
	}
}

func (s *RatingsAve) GetByPostId(postId int) (*RatingsAve, error) {
	ratingsAve := &RatingsAve{}
	query := `SELECT * FROM ratings_ave WHERE post_id = ?`
	err := config.Db.Unsafe().Get(ratingsAve, query, postId)
	if err != nil && err != sql.ErrNoRows {
		return ratingsAve, err
	}
	return ratingsAve, nil
}

func (s *RatingsAve) Upsert() error {
	query := `INSERT INTO ratings_ave (post_id, rating_ave, rating_count) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE rating_ave = VALUES(rating_ave), rating_count = VALUES(rating_count)`
	_, err := config.Db.Exec(query, s.PostId, s.RatingAve, s.RatingCount)
	if err != nil {
		return err
	}
	s.PostId = int(s.PostId)
	return nil
}
