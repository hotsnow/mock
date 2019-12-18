package app

import (
	"fmt"
	"github.com/hotsnow/api"
)

func main() {
	j := job{a: &api.NUM{}}
	d := j.task(3)
	fmt.Println(3, d)
}

type job struct {
	a api.Action
}

// double me
func (j job) task(i int) int {
	j.a.Double(&i)

	return i
}

//// Getmyip
//func (j job) Getmyip(url string) string {
//	var ip string
//
//	api.Getweb(j.c, url, &ip)
//	fmt.Println(ip)
//
//	return ip
//}
