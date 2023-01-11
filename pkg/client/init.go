package client

import (
	"github.com/wachayathorn/golang-basic-structure/pkg/utils"
)

type ClientAPI interface {
	ReplaceUrl(url string)
	GetBooks() (string, error)
}

type Client struct {
	Url   string
	Utils utils.UtilsInterface
}

func Init() ClientAPI {
	return &Client{
		Url:   "http://localhost:8080",
		Utils: utils.Init(),
	}
}

func (c *Client) ReplaceUrl(url string) {
	c.Url = url
}
