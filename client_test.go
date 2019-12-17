package app

import (
	//"fmt"
	"github.com/golang/mock/gomock"
	"github.com/hotsnow/api"
	"github.com/hotsnow/mocks"
	"testing"
)

func TestReq(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockClient(ctrl)

	url := "http://ifconfig.co"

	m.
		EXPECT().
		Req(gomock.Eq(url)).
		Return("1.2.3.4").
		AnyTimes()
		//Req(gomock.Eq(url)).

	// failed here
	j := job{c: m}

	got := j.Getmyip(url)
	if got != "1.2.3.4" {
		t.Errorf("got = %#v; want 101", got)
	}
	t.Log("Getmyip success")

	// success here
	got = api.Getweb(m, url)
	if got != "1.2.3.4" {
		t.Errorf("got = %#v; want 101", got)
	}
	t.Log("Getweb success")
}
