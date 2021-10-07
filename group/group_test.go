package group

import "testing"

func TestGetGroup(t *testing.T) {
	content, err := GetGroup("beijingzufang")
	if content != nil {
		t.Log(content.Text())
	}
	if err != nil {
		t.Fatal(err)
	}
}
