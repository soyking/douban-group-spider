package group

import "testing"

func TestGetGroup(t *testing.T) {
	content, err := GetGroup("beijingzufang")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(content.Text())
	}
}
