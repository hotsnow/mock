package app

import (
	"fmt"
	"github.com/hotsnow/api"
)

func main() {
	j := job{c: &api.WWW{}}
	j.Getmyip("http://ifconfig.co")
}

type job struct {
	c api.Client
}

// Getmyip
func (j job) Getmyip(url string) string {
	var ip string

	api.Getweb(j.c, url, &ip)
	fmt.Println(ip)

	return ip
}
