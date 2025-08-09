package controller

import (
	"encoding/json"
	"io"
	"test/store"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestListTags(t *testing.T) {
	g.POST("/test", ListTags)

	rsp := apitest.New().Handler(g).
		Post("/test").
		Expect(t).End().Response
	buf, err := io.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		t.Error(err)
		return
	}

	var rspData struct {
		List []*store.Tag `json:"list"`
	}
	if err := json.Unmarshal(buf, &rspData); err != nil {
		t.Error(err)
		return
	}
	if len(rspData.List) == 0 {
		t.Log("no tags")
		return
	}

	for _, tag := range rspData.List {
		t.Logf("%+v", tag)
	}
}
