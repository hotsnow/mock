package app

import (
	//"fmt"
	"github.com/golang/mock/gomock"
	"github.com/hotsnow/api"
	"testing"
)

func TestReq(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := api.NewMockClient(ctrl)

	url := "http://ifconfig.co"

	m.
		EXPECT().
		Req(gomock.Eq(url)).
		Return("1.2.3.4")

	got := Getmyip(url)

	if got != "1.2.3.4" {
		t.Errorf("got = %#v; want 101", got)
	}
}
