package app

import (
	"github.com/golang/mock/gomock"
	"github.com/hotsnow/mocks"
	"testing"
)

func TestReq(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAction(ctrl)
	m.EXPECT().Double(gomock.Any()).Return(nil)

	j := job{a: m}
	got := j.task(3)
	if got != 6 {
		t.Errorf("got = %#v; want 6", got)
	}
}
