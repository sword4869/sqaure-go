package store

import "testing"

func TestRatingsAve_Upsert(t *testing.T) {
	ratingsAve := &RatingsAve{
		PostId:    1,
		RatingAve: 5.1,
	}
	if err := ratingsAve.Upsert(); err != nil {
		t.Errorf("Upsert ratingsAve failed: %v", err)
		return
	}
	t.Logf("%+v", ratingsAve)
}

func TestRatingsAve_GetByPostId(t *testing.T) {
	ratingsAve, err := NewRatingsAve().GetByPostId(1)
	if err != nil {
		t.Errorf("Get ratingsAve failed: %v", err)
		return
	}
	t.Logf("%+v", ratingsAve)
}
