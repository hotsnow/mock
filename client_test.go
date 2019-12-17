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
		Req(gomock.Any()).
		Return("1.2.3.4")
		//Req(gomock.Eq(url)).

	//got := Getmyip(url)

	w := &api.WWW{}
	got := w.Req(url)
	if got != "1.2.3.4" {
		t.Errorf("got = %#v; want 101", got)
	}
}
