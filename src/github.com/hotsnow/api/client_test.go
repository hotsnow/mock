package api

import (
	//"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestReq(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockClient(ctrl)

	url := "http://ifconfig.co"

	m.
		EXPECT().
		Req(gomock.Any()).
		Return("1.2.3.4")
		//Req(gomock.Eq(url)).

	//got := Getmyip(url)

	w := &WWW{}

	got := c.Req(url)
	if got != "1.2.3.4" {
		t.Errorf("got = %#v; want 101", got)
	}

	got = Getweb(m, url)
	if got != "1.2.3.4" {
		t.Errorf("got = %#v; want 101", got)
	}
}
