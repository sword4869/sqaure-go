package controller

import (
	"io"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestPostLikeGet(t *testing.T) {
	data := `{
		"post_id": 1,
		"user_id": 1
	}`
	g.POST("/test", PostLikeGet)
	rsp := apitest.New().Handler(g).
		Post("/test").
		Body(data).
		Header("Content-Type", "application/json").
		Expect(t).End().Response
	buf, err := io.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", string(buf))
}

func TestPostLikeUpdateLiked(t *testing.T) {
	data := `{
		"post_id": 1,
		"user_id": 1,
		"is_like": 0
	}`
	g.POST("/test", PostLikeUpdateLiked)
	rsp := apitest.New().Handler(g).
		Post("/test").
		Body(data).
		Header("Content-Type", "application/json").
		Expect(t).End().Response
	buf, err := io.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", string(buf))
}
