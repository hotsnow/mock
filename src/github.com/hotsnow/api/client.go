package api

import (
	//"fmt"
	"io/ioutil"
	"net/http"
)

// Client --
type Client interface {
	Req(url string, ip *string) error
}

// WWW --
type WWW struct {
}

// Req --
func (w *WWW) Req(url string, ip *string) error {
	resp, _ := http.Get(url)

	body, _ := ioutil.ReadAll(resp.Body)

	*ip = string(body)

	return nil
}

// Getweb --
func Getweb(c Client, url string, ip *string) error {
	c.Req(url, ip)

	return nil
}
