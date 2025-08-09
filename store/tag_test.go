package store

import (
	"testing"
)

func TestTag_GetByID(t *testing.T) {
	tag := &Tag{}
	tag, err := tag.GetByID(1)
	if err != nil {
		t.Errorf("Get tag by id failed: %v", err)
		return
	}
	t.Logf("%+v", tag)
}

func TestTag_ListAllTags(t *testing.T) {
	tag := &Tag{}
	tags, err := tag.ListAllTags()
	if err != nil {
		t.Errorf("List all tags failed: %v", err)
		return
	}
	for _, tag := range tags {
		t.Logf("%+v", tag)
	}
}
