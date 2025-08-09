package store

import (
	"testing"
	"time"
)

func TestPost(t *testing.T) {
	post := NewPost()
	post.UserId = 1
	post.PostTags = []int{1, 2, 3}
	post.Content = "这是一条测试发帖，说说菜品、体验感、服务怎么样，给大家参考~"
	post.Images = []string{"1.jpg", "2.jpg"}
	post.Latitude = 120.123456
	post.Longitude = 30.123456
	post.IsActive = 1
	err := post.Create()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", post)
}

func TestPostGetByID(t *testing.T) {
	post := NewPost()
	post, err := post.GetById(1)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", post)
}

func TestPostListPosts(t *testing.T) {
	post := NewPost()
	posts, err := post.ListPosts(time.Now().Unix())
	if err != nil {
		t.Error(err)
		return
	}
	for _, post := range posts {
		t.Logf("%+v", post)
	}
}

func TestPostListPostsByKeyword(t *testing.T) {
	post := NewPost()
	posts, err := post.ListPostsByKeyword("", time.Now().Unix())
	if err != nil {
		t.Error(err)
		return
	}
	for _, post := range posts {
		t.Logf("%+v", post)
	}
}
