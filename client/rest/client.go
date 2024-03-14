package rest

import (
	"github.com/infraboard/mcube/client/rest"
)

type Client struct {
	conf *Config
	c    *rest.RESTClient
}

func NewClient(conf *Config) *Client {
	c := rest.NewRESTClient()
	c.SetBaseURL(conf.URL)
	return &Client{
		c:    c,
		conf: conf,
	}
}
