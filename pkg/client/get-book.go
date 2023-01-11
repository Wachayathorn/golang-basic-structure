package client

import (
	"net/http"

	"github.com/wachayathorn/golang-basic-structure/pkg/handler/book/httpinfo"
)

func (c *Client) GetBooks() (string, error) {
	apiPath := httpinfo.GroupV1 + httpinfo.GetBooks
	res, err := c.Utils.DoRequest(c.Url, http.MethodGet, apiPath, nil)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
