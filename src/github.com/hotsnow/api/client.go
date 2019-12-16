package api

import (
	//"fmt"
	"io/ioutil"
	"net/http"
)

// Client --
type Client interface {
	Req(url string) string
}

// WWW --
type WWW struct {
}

// Req --
func (w *WWW) Req(url string) string {
	resp, _ := http.Get(url)

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

// Getweb --
func Getweb(c Client, url string) string {
	return c.Req(url)
}
