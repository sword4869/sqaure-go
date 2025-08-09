package store

import (
	"testing"
)

func TestRatings_Create(t *testing.T) {
	ratings := &Ratings{
		UserId: 1,
		PostId: 1,
		Rating: 5,
	}
	if err := ratings.Create(); err != nil {
		t.Errorf("Create ratings failed: %v", err)
		return
	}
	t.Logf("%+v", ratings)
}

func TestRatings_GetByUserIdAndPostId(t *testing.T) {
	ratings, err := NewRatings().GetByUserIdAndPostId(1, 1)
	if err != nil {
		t.Errorf("Get ratings failed: %v", err)
		return
	}
	t.Logf("%+v", ratings)
}
