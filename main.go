package app

import (
	//"fmt"
	"github.com/hotsnow/api"
)

//func main() {
//	myip("http://ifconfig.co")
//}

// Getmyip
func Getmyip(url string) string {
	w := api.WWW{}
	return api.Getweb(&w, url)
}
