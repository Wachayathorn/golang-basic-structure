package client

import (
	"encoding/json"
	"net/http"

	bookhttpinfo "github.com/wachayathorn/golang-basic-structure/pkg/handler/book/httpinfo"
	bookrepository "github.com/wachayathorn/golang-basic-structure/pkg/repository/book"
)

func (c *Client) GetBooks() ([]bookrepository.Book, error) {
	apiPath := bookhttpinfo.GroupV1 + bookhttpinfo.GetBooks
	res, err := c.Utils.DoRequest(c.Url, http.MethodGet, apiPath, nil)
	if err != nil {
		return nil, err
	}

	var result []bookrepository.Book
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}
	return result, nil
}
