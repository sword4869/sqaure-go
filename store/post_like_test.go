package store

import (
	"testing"
	"time"
)

func TestPostLike_Create(t *testing.T) {
	postLike := &PostLike{
		UserId:    1,
		PostId:    1,
		IsLike:    1,
		CreatedAt: time.Now().Unix(),
	}
	if err := postLike.Create(); err != nil {
		t.Errorf("Create post like failed: %v", err)
		return
	}
	t.Logf("%+v", postLike)
}

func TestPostLike_GetByUserIdAndPostId(t *testing.T) {
	postLike := &PostLike{
		UserId: 1,
		PostId: 1,
	}
	postLike, err := postLike.GetByUserIdAndPostId(1, 1)
	if err != nil {
		t.Errorf("Get post like failed: %v", err)
		return
	}
	t.Logf("%+v", postLike)
}

func TestPostLike_UpdateLiked(t *testing.T) {
	postLike := &PostLike{
		UserId: 1,
		PostId: 1,
		IsLike: 0,
	}
	if err := postLike.UpdateLiked(1, 1, 0); err != nil {
		t.Errorf("Update post like failed: %v", err)
		return
	}
	t.Logf("%+v", postLike)
}
