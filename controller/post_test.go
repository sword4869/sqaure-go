package controller

import (
	"encoding/json"
	"io"
	"test/store"
	"testing"

	"github.com/steinfletcher/apitest"
)

// TestCreatePost 测试创建发帖（使用gin测试）
func TestCreatePost(t *testing.T) {
	data := `{
		"user_id": 1,
		"post_tag": 1,
		"content": "这是一条测试发帖，说说菜品、体验感、服务怎么样，给大家参考~",
		"images": ["https://example.com/image1.jpg", "https://example.com/image2.jpg"],
		"latitude": 39.9042,
		"longitude": 116.4074
	}`
	g.POST("/test", CreatePost)
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
	t.Log(string(buf))
}

func TestListPosts(t *testing.T) {
	data := `{
		"cursor": 0
	}`
	g.POST("/test", ListPosts)
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

func TestListPostsByKeyword(t *testing.T) {
	data := `{
		"cursor": 0,
		"keyword": "测试"
	}`
	g.POST("/test", ListPostsByKeyword)
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
	posts := make([]*store.Post, 0)
	json.Unmarshal(buf, &posts)
	for _, post := range posts {
		t.Logf("%+v", post)
	}
}
